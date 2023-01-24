package operations

type GetAPITagsIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAPITagsID200ApplicationJSONOrders struct {
	ID *int64 `json:"id,omitempty"`
}

type GetAPITagsID200ApplicationJSON struct {
	Name   *string                                `json:"name,omitempty"`
	Orders []GetAPITagsID200ApplicationJSONOrders `json:"orders,omitempty"`
}

type GetAPITagsIDRequest struct {
	PathParams GetAPITagsIDPathParams
}

type GetAPITagsIDResponse struct {
	ContentType                          string
	StatusCode                           int64
	GetAPITagsID200ApplicationJSONObject *GetAPITagsID200ApplicationJSON
}
