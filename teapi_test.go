package teapi

import (
	"bytes"
	"compress/gzip"
	. "github.com/karlseguin/expect"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func init() {
	BeforeEach(func() {
		setupFakeNow(time.Unix(872838840, 0))
		stubDoOK()
	})
}

type TeapiTests struct{}

func Test_Teapi(t *testing.T) {
	Expectify(new(TeapiTests), t)
}

func (_ TeapiTests) RetriesTimestampFailures() {
	stubDo(401, `{"error": "invalid timestamp", "date": "server-date"}`, nil)
	n().Request("POST", "documents", bytes.NewReader([]byte{}))
	Expect(len(requests)).To.Equal(2)
	last := requests[len(requests)-1]
	Expect(last.Header.Get("Date")).To.Equal("server-date")
}

func setupFakeNow(now time.Time) {
	Now = func() time.Time {
		return now
	}
}

var bodies [][]byte
var requests []*http.Request

func stubDo(code int, body string, err error) {
	res := &http.Response{
		StatusCode: code,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(body))),
	}
	bodies = make([][]byte, 0, 2)
	requests = make([]*http.Request, 0, 2)
	Do = func(c *http.Client, req *http.Request) (*http.Response, error) {
		requests = append(requests, req)
		b, _ := ioutil.ReadAll(req.Body)
		bodies = append(bodies, b)
		return res, err
	}
}

func stubDoOK() {
	stubDo(200, "", nil)
}

func assertLast(method, url, expected, signature string) {
	assertLastHeaders(method, url, signature)
	body := bodies[len(bodies)-1]
	Expect(body).To.Equal(JSON(expected))
}

func assertLastGzip(method, url, expected, signature string) {
	assertLastHeaders(method, url, signature)
	buffer := bytes.NewBuffer(bodies[len(bodies)-1])
	gr, err := gzip.NewReader(buffer)
	if err != nil {
		panic(err)
	}
	defer gr.Close()
	body, err := ioutil.ReadAll(gr)
	if err != nil {
		panic(err)
	}
	Expect(body).To.Equal(JSON(expected))
}

func assertLastHeaders(method, url, signature string) {
	last := requests[len(requests)-1]
	Expect(last.Method).To.Equal(method)
	Expect(last.URL.String()).To.Equal(url)
	Expect(last.Header.Get("Date")).To.Equal("Fri, 29 Aug 1997 14:14:00 GMT")
	Expect(last.Header.Get("Authorization")).To.Equal("HMAC-SHA256 Credential=leto,Signature=" + signature)
}

func n() *Teapi {
	return New(Configure("test.teapi.io", "leto", "over9000!"))
}

type P struct {
	Id   int    `json:"id",omitempty`
	Name string `json:"name",omitempty`
}

func Person(id int, name string) *P {
	return &P{id, name}
}
