package teapi

type lists struct {
	*Teapi
}

// Insert an document id into the list
func (l lists) Insert(t string, list string, ids ...interface{}) (int, error) {
	var data = map[string]interface{}{"type": t, "list": list, "ids": ids, "truncate": false}
	return l.request("POST", "lists", data)
}

// Replaces the existing list ids with the supplied ids (truncate+insert)
func (l lists) Replace(t string, list string, ids ...interface{}) (int, error) {
	var data = map[string]interface{}{"type": t, "list": list, "ids": ids, "truncate": true}
	return l.request("POST", "lists", data)
}

// Delete the ids from the list
func (l lists) Delete(t string, list string, ids ...interface{}) (int, error) {
	var data = map[string]interface{}{"type": t, "list": list, "ids": ids}
	return l.request("DELETE", "lists", data)
}
