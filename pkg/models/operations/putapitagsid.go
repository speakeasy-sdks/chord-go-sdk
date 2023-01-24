package operations

type PutAPITagsIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type PutAPITagsIDRequestBodyTag struct {
	Name string `json:"name"`
}

type PutAPITagsIDRequestBody struct {
	Tag PutAPITagsIDRequestBodyTag `json:"tag"`
}

type PutAPITagsID200ApplicationJSONOrders struct {
	ID *int64 `json:"id,omitempty"`
}

type PutAPITagsID200ApplicationJSON struct {
	Name   *string                                `json:"name,omitempty"`
	Orders []PutAPITagsID200ApplicationJSONOrders `json:"orders,omitempty"`
}

type PutAPITagsIDRequest struct {
	PathParams PutAPITagsIDPathParams
	Request    *PutAPITagsIDRequestBody `request:"mediaType=application/json"`
}

type PutAPITagsIDResponse struct {
	ContentType                          string
	StatusCode                           int64
	PutAPITagsID200ApplicationJSONObject *PutAPITagsID200ApplicationJSON
}
