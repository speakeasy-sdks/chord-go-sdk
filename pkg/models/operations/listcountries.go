package operations

import (
	"net/http"
)

type ListCountriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
