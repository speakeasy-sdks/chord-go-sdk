package operations

import (
	"net/http"
)

type PutAPIRolesIDRemoveUserIDPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type PutAPIRolesIDRemoveUserIDRequest struct {
	PathParams PutAPIRolesIDRemoveUserIDPathParams
}

type PutAPIRolesIDRemoveUserIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
