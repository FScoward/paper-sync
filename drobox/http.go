package drobox

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO: Error handle
func GetDocIdList(client *http.Client) []byte {
	jsonBuf := bytes.NewBuffer([]byte("{}"))
	req, _ := http.NewRequest(http.MethodPost, DocList, jsonBuf)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := post(req, client)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func DownloadDoc(client *http.Client, docId string, format string) (DownloadDocResponse, error) {
	req, _ := http.NewRequest(http.MethodPost, DownLoad, nil)
	arg := fmt.Sprintf(`{"doc_id":"%s","export_format":{".tag":"%s"}}`, docId, format)
	req.Header.Set("Dropbox-API-Arg", arg)
	resp, _ := post(req, client)

	defer resp.Body.Close()
	downloadDocResponse := DownloadDocResponse{}

	return downloadDocResponse.From(resp)
}

func UpdateDoc(client *http.Client, filePath string, docId string, policy string, revision int, format string) (*http.Response, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	read, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	body.Write(read)
	req, _ := http.NewRequest(http.MethodPost, Update, body)
	arg := fmt.Sprintf(`{"doc_id": "%s", "doc_update_policy": {".tag":"%s"}, "revision": %d, "import_format": "%s"}`, docId, policy, revision, format)
	req.Header.Set("Dropbox-API-Arg", arg)
	req.Header.Set("Content-Type", "application/octet-stream")

	return post(req, client)
}

func post(req *http.Request, client *http.Client) (*http.Response, error) {
	accessKey := os.Getenv("DROPBOX_PAPER_ACCESS_TOKEN")
	req.Header.Set("Authorization", "Bearer "+accessKey)
	resp, err := client.Do(req)
	return resp, err
}
