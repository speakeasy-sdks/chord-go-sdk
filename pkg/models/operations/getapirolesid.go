package operations

type GetAPIRolesIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAPIRolesIDRequest struct {
	PathParams GetAPIRolesIDPathParams
}

type GetAPIRolesID200ApplicationJSONUsers struct {
	Email *string `json:"email,omitempty"`
	ID    *int64  `json:"id,omitempty"`
}

type GetAPIRolesID200ApplicationJSON struct {
	Name  *string                                `json:"name,omitempty"`
	Users []GetAPIRolesID200ApplicationJSONUsers `json:"users,omitempty"`
}

type GetAPIRolesIDResponse struct {
	ContentType                           string
	StatusCode                            int
	GetAPIRolesID200ApplicationJSONObject *GetAPIRolesID200ApplicationJSON
}
