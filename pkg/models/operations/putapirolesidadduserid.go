package operations

import (
	"net/http"
)

type PutAPIRolesIDAddUserIDPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type PutAPIRolesIDAddUserIDRequest struct {
	PathParams PutAPIRolesIDAddUserIDPathParams
}

type PutAPIRolesIDAddUserIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
