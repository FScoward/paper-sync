package drobox

import "fmt"

type UpdateDocResponse struct {
	DocId    string `json:"doc_id"`
	Revision int    `json:"revision"`
	Title    string `json:"title"`
}

type UpdateDocErrorResponse struct {
	ErrorSummary string `json:"error_summary"`
	ErrorInfo    struct {
		Tag string `json:".tag"`
	} `json:"error"`
}

func (e *UpdateDocErrorResponse) Error() string {
	return fmt.Sprintf("error summary: %s, error: %s", e.ErrorSummary, e.ErrorInfo)
}
