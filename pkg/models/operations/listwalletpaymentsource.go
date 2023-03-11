package operations

import (
	"net/http"
)

type ListWalletPaymentSourceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
