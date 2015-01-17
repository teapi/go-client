# teapi.io Golang Library
This is a node driver for <http://teapi.io>.

# Installation

```go
go get github.com/teapi/go-client
```

# Configuration
Create a teapi instance via `New` and configure with `Configure`:

```go
import (
  "github.com/teapi/go-client"
)
...
t := teapi.New(teapi.Configure("m01.teapi.io", "KEY", "SECRET"))
```

# Usage
Documents can be created, updated or deleted one at a time:

```go
type Saiyan struct {
  Id int       `json:"id"`
  Name string  `json:"name"`
  Power int    `json:"power"`
}

user := &Saiyan{Id: 34, Name: "Goku", Power: 9001}

t.Documents.Create("saiyans", teapi.Doc(user))
t.Documents.Update("saiyans", teapi.Doc(user))
t.Documents.Delete("saiyans", 34)
```

`teapi.DocMeta` can be used instead of `teapi.Doc`. The meta object is used to provide index values that shouldn't be returned as part of the main document:

```go
override := map[string]interface{"power", 49943, "super": true}
t.Documents.Create("saiyans", teapi.DocMeta(doc, override))
```
`Doc` and `DocMeta` support any type that can be serialized via `encoding/json`.

# Bulk Documents
Documents can be inserted, updated and deleted in bulk:

```go
upserts := teapi.Documents {
  teapi.Doc(user1), teapi.Doc(user2),
}
deletes := teapi.DocumentIds{teapi.DocId(434), teapi.DocId("string_id"),}

t.Documents.Bulk("saiyans", upserts, deletes)
```

Up to 1000 items can be sent per call.

## Lists
Document ids can be pushed onto and removed from lists:

```go
t.Lists.Insert('people', 'newest', 1, 4, 992)
t.Lists.Delete('people', 'newest', 992)
```

For both `Insert` and `Delete`, the ids can either be a variable number of ids (up to 1000 ids can be inserted / deleted at a time). You can also pass no ids to `Delete`. This will remove all ids from the list:

```go
t.Lists.Delete('people', 'newest')
```

`Replace` is similar to `Insert`, but any existing values in the list will first be removed (it truncates the existing list and inserts the new values):

```go
// this will result in a list of only 1 id (99)
t.Lists.Replace('people', 'newest', 99,)
```

# Return value
The return value is an `(int, error)` where the integer represents the status code of the response. A negative status code is possible if the request wasn't possible (such as on a serialization error). It's also possible to get both status code and an error (such as a 401 authorization error).

# Versioning
The library is available via `gopkg.in`. Install and import `gopkg.in/teapi/teapi.v1` rather than the `github` variant.
