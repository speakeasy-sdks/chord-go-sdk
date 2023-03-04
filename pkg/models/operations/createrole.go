package operations

import (
	"net/http"
)

type CreateRoleRequestBodyRoleRoleUsersAttributes struct {
	UserID *int64 `json:"user_id,omitempty"`
}

type CreateRoleRequestBodyRole struct {
	Name                string                                         `json:"name"`
	RoleUsersAttributes []CreateRoleRequestBodyRoleRoleUsersAttributes `json:"role_users_attributes,omitempty"`
}

type CreateRoleRequestBody struct {
	Role CreateRoleRequestBodyRole `json:"role"`
}

type CreateRoleRequest struct {
	Request *CreateRoleRequestBody `request:"mediaType=application/json"`
}

type CreateRoleResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
