package drobox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestDownloadDoc(t *testing.T) {
	t.Run("factory method success", func(t *testing.T) {

		owner := "fscoward"
		title := "download markdown"
		revision := 100
		mimeType := "text/x-markdown"
		body := "test"

		jsonString := fmt.Sprintf(`{
			"owner": "%s",
			"title": "%s",
			"revision": %d,
			"mime_type": "%s" 
			}`, owner, title, revision, mimeType)

		header := map[string][]string{
			DROPBOX_API_RESULT_KEY: {jsonString}}

		response := httptest.ResponseRecorder{
			Code:      200,
			HeaderMap: header,
			Body:      bytes.NewBufferString(body),
			Flushed:   false,
		}

		factory := DownloadDocResponse{}
		actual := factory.From(response.Result())

		expect := DownloadDocResponse{
			Header: DownloadDocHeader{
				Owner:    owner,
				Title:    title,
				Revision: revision,
				MimeType: mimeType,
			},
			Body: body,
		}

		if actual != expect {
			t.Errorf("Unexpected error. actual: %v, expect: %v", actual, expect)
		}
	})

	t.Run("json parse success", func(t *testing.T) {

		jsonString := `{
			"owner": "fscoward",
			"title": "download markdown",
			"revision": 100,
			"mime_type": "text/x-markdown"
			}`

		jsonBuf := []byte(jsonString)

		var downloadDoc DownloadDocHeader

		err := json.Unmarshal(jsonBuf, &downloadDoc)
		if err != nil {
			fmt.Println(err)
		}

		expect := DownloadDocHeader{
			"fscoward",
			"download markdown",
			100,
			"text/x-markdown"}

		if expect != downloadDoc {
			t.Errorf("unexpected actual: %v, expect: %v", downloadDoc, expect)
		}
	})

}
