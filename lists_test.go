package teapi

import (
	. "github.com/karlseguin/expect"
	"testing"
)

type ListsTests struct{}

func Test_Lists(t *testing.T) {
	Expectify(new(ListsTests), t)
}

func (_ ListsTests) InsertsOneId() {
	Expect(n().Lists.Insert("atreides", "newest", "leto2")).To.Equal(200, nil)
	assertLast(
		"POST",
		"https://test.teapi.io/v1/lists",
		`{"type":"atreides", "list": "newest", "ids": ["leto2"], "truncate": false}`,
		"dbfdfcd90e7b5124e4807d156200fa19547359215719c93ccbadd43d164e110f")
}

func (_ ListsTests) InsertsManyIds() {
	Expect(n().Lists.Insert("atreides", "newest", 1, 10, 99)).To.Equal(200, nil)
	assertLast(
		"POST",
		"https://test.teapi.io/v1/lists",
		`{"type":"atreides", "list": "newest", "ids": [1,10,99], "truncate": false}`,
		"f9eb0175d9a62d139550910169699437d7e3c650cdd211ec68453aeb4b107bd9")
}

func (_ ListsTests) ReplacesListWithOneId() {
	Expect(n().Lists.Replace("atreides", "newest", 442)).To.Equal(200, nil)
	assertLast(
		"POST",
		"https://test.teapi.io/v1/lists",
		`{"type":"atreides", "list": "newest", "ids": [442], "truncate": true}`,
		"b39bae65b9efee0390df05786dfb7d93b32d1b44b1b107385cc39d8102493de1")
}

func (_ ListsTests) ReplacesListWithManyIds() {
	Expect(n().Lists.Replace("atreides", "newest", "1a", "2a")).To.Equal(200, nil)
	assertLast(
		"POST",
		"https://test.teapi.io/v1/lists",
		`{"type":"atreides", "list": "newest", "ids": ["1a", "2a"], "truncate": true}`,
		"69258a3d959f39ce5b8bee35b241c458a70ab324accff5ee3aac7f3b02b73788")
}

func (_ ListsTests) DeletesOneId() {
	Expect(n().Lists.Delete("atreides", "newest", "kakk40")).To.Equal(200, nil)
	assertLast(
		"DELETE",
		"https://test.teapi.io/v1/lists",
		`{"type":"atreides", "list": "newest", "ids": ["kakk40"]}`,
		"1c5bc05e998ac822a7a52b20ac960a1fecf8bcd2a61a0e1ffb7fa5e699c4bc4a")
}

func (_ ListsTests) DeletesManyIds() {
	Expect(n().Lists.Delete("atreides", "newest", 1, 55, 209)).To.Equal(200, nil)
	assertLast(
		"DELETE",
		"https://test.teapi.io/v1/lists",
		`{"type":"atreides", "list": "newest", "ids": [1,55,209]}`,
		"de2904edb6733a7203691ac6c00cc0504997ad33a793027fe2be9d1fd94a0a7f")
}

func (_ ListsTests) DeletesAll() {
	Expect(n().Lists.Delete("atreides", "newest")).To.Equal(200, nil)
	assertLast(
		"DELETE",
		"https://test.teapi.io/v1/lists",
		`{"type":"atreides", "list": "newest", "ids": null}`,
		"beb5856df4540fa5aeeb8f1879f188edcaa7798a47dbaf6a87600800a746c3ab")
}
