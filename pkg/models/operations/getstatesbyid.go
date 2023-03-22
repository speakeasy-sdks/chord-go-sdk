// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetStatesByIDPathParams struct {
	CountryID string `pathParam:"style=simple,explode=false,name=country_id"`
}

type GetStatesByIDRequest struct {
	PathParams GetStatesByIDPathParams
}

type GetStatesByIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
