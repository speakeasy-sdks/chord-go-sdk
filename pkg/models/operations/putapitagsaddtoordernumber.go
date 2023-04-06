// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPITagsAddToOrderNumberRequestBody struct {
	Tags []string `json:"tags,omitempty"`
}

type PutAPITagsAddToOrderNumberRequest struct {
	RequestBody *PutAPITagsAddToOrderNumberRequestBody `request:"mediaType=application/json"`
	OrderNumber string                                 `pathParam:"style=simple,explode=false,name=order_number"`
}

type PutAPITagsAddToOrderNumberResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
