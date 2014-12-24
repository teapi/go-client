package teapi

import (
	"bytes"
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
	Expect(last.Header.Get("Date")).To.Equal("server-date")
}

func setupFakeNow(now time.Time) {
	Now = func() time.Time {
		return now
	}
}

var last *http.Request
var requests []*http.Request

func stubDo(code int, body string, err error) {
	res := &http.Response{
		StatusCode: code,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(body))),
	}
	requests = make([]*http.Request, 0, 2)
	Do = func(c *http.Client, req *http.Request) (*http.Response, error) {
		last = req
		requests = append(requests, req)
		return res, err
	}
}

func stubDoOK() {
	stubDo(200, "", nil)
}

func assertLast(method, url, expected, signature string) {
	Expect(last.Method).To.Equal(method)
	Expect(last.URL.String()).To.Equal(url)
	Expect(last.Header.Get("Date")).To.Equal("Fri, 29 Aug 1997 14:14:00 GMT")
	Expect(last.Header.Get("Authorization")).To.Equal("HMAC-SHA256 Credential=leto,Signature=" + signature)
	actual, _ := ioutil.ReadAll(last.Body)
	Expect(actual).To.Equal(JSON(expected))
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
