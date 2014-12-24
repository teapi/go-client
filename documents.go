package teapi

type documents struct {
	*Teapi
}

// Create a new document
func (d documents) Create(t string, doc *Document) (int, error) {
	var data = map[string]interface{}{"type": t, "doc": doc.Doc}
	if doc.Meta != nil {
		data["meta"] = doc.Meta
	}
	return d.request("POST", "documents", data)
}

// Updates an existing document
func (d documents) Update(t string, doc *Document) (int, error) {
	var data = map[string]interface{}{"type": t, "doc": doc.Doc}
	if doc.Meta != nil {
		data["meta"] = doc.Meta
	}
	return d.request("PUT", "documents", data)
}

// Updates an existing document
func (d documents) Delete(t string, id interface{}) (int, error) {
	var data = map[string]interface{}{"type": t, "id": id}
	return d.request("DELETE", "documents", data)
}

func (d documents) Bulk(t string, upserts Documents, deletes DocumentIds) (int, error) {
	var data = map[string]interface{}{"type": t, "upserts": upserts, "deletes": deletes}
	return d.request("POST", "documents", data)
}

type Document struct {
	Doc  interface{} `json:"doc"`
	Meta interface{} `json:"meta",omitempty`
}

type DocumentId struct {
	Id interface{} `json:"id"`
}

func Doc(doc interface{}) *Document {
	return &Document{doc, nil}
}

func DocMeta(doc, meta interface{}) *Document {
	return &Document{doc, meta}
}

func DocId(id interface{}) *DocumentId {
	return &DocumentId{id}
}

type DocDeletes struct {
	Id interface{} `json:"id"`
}

type Documents []*Document
type DocumentIds []*DocumentId
