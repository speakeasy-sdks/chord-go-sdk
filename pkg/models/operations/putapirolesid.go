package operations

type PutAPIRolesIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type PutAPIRolesIDRequestBodyRoleUsers struct {
	Destroy *bool `json:"_destroy,omitempty"`
	ID      int64 `json:"id"`
}

type PutAPIRolesIDRequestBodyRole struct {
	Name  string                              `json:"name"`
	Users []PutAPIRolesIDRequestBodyRoleUsers `json:"users,omitempty"`
}

type PutAPIRolesIDRequestBody struct {
	Role PutAPIRolesIDRequestBodyRole `json:"role"`
}

type PutAPIRolesIDRequest struct {
	PathParams PutAPIRolesIDPathParams
	Request    *PutAPIRolesIDRequestBody `request:"mediaType=application/json"`
}

type PutAPIRolesIDResponse struct {
	ContentType string
	StatusCode  int
}
