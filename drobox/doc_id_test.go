package drobox

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDocIdListParse(t *testing.T) {
	jsonString := `{
	  "doc_ids": [
		"aaaaaaaaaaaaaaaaaaaaa",
		"bbbbbbbbbbbbbbbbbbbbb"
	  ],
	  "cursor": {
		"value": "value_sample",
		"expiration": "2000-01-01T09:00:00Z"
	  },
	  "has_more": false
	}`
	jsonBuf := []byte(jsonString)

	var docIdList DocIdList

	err := json.Unmarshal(jsonBuf, &docIdList)
	if err != nil {
		fmt.Println(err)
	}

	expect := DocIdList{
		[]string{"aaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbb"},
		cursor{"value_sample", time.Date(2000, 1, 1, 9, 0, 0, 0, time.UTC)},
		false}

	if !reflect.DeepEqual(expect.DocIds, docIdList.DocIds) {
		t.Errorf("unexpected actual: %v, expect: %v", docIdList.DocIds, expect.DocIds)
	}

	if expect.Cursor != docIdList.Cursor {
		t.Errorf("unexpected actual: %v, expect: %v", docIdList.Cursor, expect.Cursor)
	}

	if expect.HasMore != docIdList.HasMore {
		t.Errorf("unexpected actual: %v, expect: %v", docIdList.HasMore, expect.HasMore)
	}

}
