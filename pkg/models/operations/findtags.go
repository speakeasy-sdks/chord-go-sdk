package operations

import (
	"net/http"
)

type FindTagsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
