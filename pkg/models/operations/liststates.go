package operations

import (
	"net/http"
)

type ListStatesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
