package drobox

import (
	"net/http"
	"os"
	"bytes"
	"io/ioutil"
)

// TODO: Error handle
func GetDocIdList(client *http.Client) []byte {
	jsonBuf := bytes.NewBuffer([]byte("{}"))
	req, _  := http.NewRequest(http.MethodPost, DocList, jsonBuf)

	//client := new(http.Client)
	resp, _ := Post(req, client)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}

func Post(req *http.Request, client *http.Client) (*http.Response, error) {
	accessKey := os.Getenv("DROPBOX_PAPER_ACCESS_TOKEN")
	req.Header.Set("Authorization", "Bearer " + accessKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	return resp, err
}