package teapi

import (
	. "github.com/karlseguin/expect"
	"testing"
)

type DocumentsTests struct{}

func Test_Documents(t *testing.T) {
	Expectify(new(DocumentsTests), t)
}

func (_ DocumentsTests) CreatesADocument() {
	Expect(n().Documents.Create("atreides", Doc(Person(434, "Leto")))).To.Equal(200, nil)
	assertLast(
		"POST",
		"https://test.teapi.io/v1/documents",
		`{"type":"atreides", "doc":{"id":434, "name":"Leto"}}`,
		"9b8bcb7da2e2b360f7db2be82241b6355d5068a55bc1bd7134887e0d3dc35d84")
}

func (_ DocumentsTests) CreatesADocumentWithMeta() {
	Expect(n().Documents.Create("atreides", DocMeta(Person(434, "Leto"), Person(22, "")))).To.Equal(200, nil)
	assertLast(
		"POST",
		"https://test.teapi.io/v1/documents",
		`{"type":"atreides", "doc":{"id":434, "name":"Leto"}, "meta":{"id": 22, "name":""}}`,
		"d8c84e4eb409ca176299159434923e1ff54cc30151dc8b3535b80eff79cc2148")
}

func (_ DocumentsTests) UpdatesADocument() {
	Expect(n().Documents.Update("atreides", Doc(Person(434, "Leto")))).To.Equal(200, nil)
	assertLast(
		"PUT",
		"https://test.teapi.io/v1/documents",
		`{"type":"atreides", "doc":{"id":434, "name":"Leto"}}`,
		"9b8bcb7da2e2b360f7db2be82241b6355d5068a55bc1bd7134887e0d3dc35d84")
}

func (_ DocumentsTests) UpdatesADocumentWithMeta() {
	Expect(n().Documents.Update("saiyans", DocMeta(Person(9000, "Goku"), Person(9001, "")))).To.Equal(200, nil)
	assertLast(
		"PUT",
		"https://test.teapi.io/v1/documents",
		`{"type":"saiyans", "doc":{"id":9000, "name":"Goku"}, "meta":{"id": 9001, "name":""}}`,
		"b4dcbfea837224cd9bf3109d9f02cd6fe46e2935bd333bf1991ec433a9db85e1")
}

func (_ DocumentsTests) DeletesADocument() {
	Expect(n().Documents.Delete("saiyans", 95323)).To.Equal(200, nil)
	assertLast(
		"DELETE",
		"https://test.teapi.io/v1/documents",
		`{"type":"saiyans", "id":95323}`,
		"85f235c49853f75e8982e0e6bb8b24bfa14a9ea2b7a34f885aac3e755ac1a665")
}

func (_ DocumentsTests) DeletesADocumentWithStringId() {
	Expect(n().Documents.Delete("saiyans", "vegeta")).To.Equal(200, nil)
	assertLast(
		"DELETE",
		"https://test.teapi.io/v1/documents",
		`{"type":"saiyans", "id":"vegeta"}`,
		"9daabeeb5c24def5aeee1fd94492770a3be432323a701a65a547ddd497f978c3")
}
