package drobox

import "time"

type DocIdList struct {
	DocIds []string  `json:"doc_ids"`
	Cursor  cursor   `json:"cursor"`
	HasMore bool     `json:"has_more"`
}

type cursor struct {
	Value string      `json:"value"`
	Expiration time.Time `json:"expiration"`
}

