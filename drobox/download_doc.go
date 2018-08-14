package drobox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const DROPBOX_API_RESULT_KEY = "Dropbox-Api-Result"

type DownloadDocHeader struct {
	Owner    string `json:"owner"`
	Title    string `json:"title"`
	Revision int    `json:"revision"`
	MimeType string `json:"mime_type"`
}

type DownloadDocResponse struct {
	Header DownloadDocHeader
	Body   string
}

func (downloadDocResponse *DownloadDocResponse) From(httpResponse *http.Response) (DownloadDocResponse, error) {
	header := httpResponse.Header
	dropboxApiResult := header.Get(DROPBOX_API_RESULT_KEY)
	var downloadHeader DownloadDocHeader
	err := json.Unmarshal([]byte(dropboxApiResult), &downloadHeader)
	if err != nil {
		// TODO
		fmt.Println(fmt.Errorf("Error occured... ", err))
	}

	body, _ := ioutil.ReadAll(httpResponse.Body)

	return DownloadDocResponse{
		downloadHeader,
		string(body),
	}, err
}

func (downloadDocResponse *DownloadDocResponse) Download() {
	ioutil.WriteFile(downloadDocResponse.Header.Title+".md", []byte(downloadDocResponse.Body), os.ModePerm)
}
