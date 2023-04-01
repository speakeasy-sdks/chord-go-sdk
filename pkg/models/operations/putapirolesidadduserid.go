// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PutAPIRolesIDAddUserIDRequest struct {
	// Role id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// User id
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type PutAPIRolesIDAddUserIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
