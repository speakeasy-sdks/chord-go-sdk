package operations

import (
	"net/http"
)

type FindRolesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
