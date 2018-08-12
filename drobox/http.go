package drobox

import (
	"net/http"
	"os"
	"bytes"
	"io/ioutil"
	"fmt"
)

// TODO: Error handle
func GetDocIdList(client *http.Client) []byte {
	jsonBuf := bytes.NewBuffer([]byte("{}"))
	req, _  := http.NewRequest(http.MethodPost, DocList, jsonBuf)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := post(req, client)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func DownloadDoc(client *http.Client, docId string, format string) []byte {
	req, _ := http.NewRequest(http.MethodPost, DownLoad, nil)
	arg := fmt.Sprintf(`{"doc_id":"%s","export_format":{".tag":"%s"}}`, docId, format)
	req.Header.Set("Dropbox-API-Arg", arg)
	resp, _ := post(req, client)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func post(req *http.Request, client *http.Client) (*http.Response, error) {
	accessKey := os.Getenv("DROPBOX_PAPER_ACCESS_TOKEN")
	req.Header.Set("Authorization", "Bearer " + accessKey)
	resp, err := client.Do(req)
	return resp, err
}