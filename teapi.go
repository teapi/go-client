// management api client for teapi.io
package teapi

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	Version = "v1"
)

type Teapi struct {
	*Configuration
	Client    *http.Client
	Documents *documents
	Lists     *lists
}

// Create a new teapi instance using the specified configuration
// Use teapi.Configure(HOST, KEY, SECRET) to create a new configuration
// object
func New(config *Configuration) *Teapi {
	t := &Teapi{
		Configuration: config,
		Client:        http.DefaultClient,
	}
	t.Documents = &documents{t}
	t.Lists = &lists{t}
	return t
}

// Send a request directly to teapi, circumventing any data encoding.
// Useful if you want to provide your own json serialization, but you'll
// have to make sure that your body follows teapi's API. For example,
// the body for create a document looks like: {"type": "TYPE", "doc": {THE_DOCUMENT}}
// (with an optional "meta" field)
func (t *Teapi) Request(method, resource string, body io.ReadSeeker) (int, error) {
	return t.do(method, resource, "", body)
}

func (t *Teapi) request(method, resource string, body interface{}) (int, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return -1, err
	}
	return t.Request(method, resource, bytes.NewReader(data))
}

func (t *Teapi) do(method, resource, date string, body io.ReadSeeker) (int, error) {
	forcedDate := true
	if len(date) == 0 {
		forcedDate = false
		date = Now().Format(http.TimeFormat)
	}

	path := "/" + Version + "/" + resource
	sig, l, err := t.sign(path, date, body)
	if err != nil {
		return -2, err
	}

	compress := l > 512
	pr, pw := io.Pipe()
	defer pr.Close()

	go func() {
		var target io.WriteCloser = pw
		var gz io.WriteCloser
		if compress {
			gz = gzip.NewWriter(pw)
			target = gz
		}
		body.Seek(0, 0)
		io.Copy(target, body)
		if compress {
			gz.Close()
		}
		pw.Close()
	}()

	req, err := http.NewRequest(method, t.host+path, pr)
	if err != nil {
		return -3, err
	}
	req.Header.Set("Date", date)
	req.Header.Set("Authorization", sig)
	if compress {
		req.Header.Set("Content-Encoding", "gzip")
	}

	code := -4
	res, err := Do(t.Client, req)

	if res != nil {
		code = res.StatusCode
	}
	if err != nil {
		return code, err
	}
	defer res.Body.Close()
	if code < 300 || code == 404 {
		return code, nil
	}

	rbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -5, errors.New(string(rbody))
	}

	var data map[string]interface{}
	if err := json.Unmarshal(rbody, &data); err != nil {
		return -6, errors.New(string(rbody))
	}

	if code == 401 && forcedDate == false {
		if date, exists := data["date"].(string); exists {
			body.Seek(0, 0)
			return t.do(method, resource, date, body)
		}
	}

	if error, exists := data["error"].(string); exists {
		return code, errors.New(error)
	}
	return -7, errors.New(string(rbody))
}

func (t *Teapi) sign(url, date string, body io.Reader) (string, int, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return "", 0, err
	}
	hasher := hmac.New(sha256.New, t.secret)
	hasher.Write([]byte(url))
	hasher.Write([]byte(date))
	hasher.Write(data)
	sig := "HMAC-SHA256 Credential=" + t.key + ",Signature=" + hex.EncodeToString(hasher.Sum(nil))
	return sig, len(data), nil
}

var Now = func() time.Time {
	return time.Now().UTC()
}

var Do = func(client *http.Client, req *http.Request) (*http.Response, error) {
	return client.Do(req)
}
