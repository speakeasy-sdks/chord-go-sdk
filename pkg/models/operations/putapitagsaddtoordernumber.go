package operations

type PutAPITagsAddToOrderNumberPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type PutAPITagsAddToOrderNumberRequestBody struct {
	Tags []string `json:"tags,omitempty"`
}

type PutAPITagsAddToOrderNumberRequest struct {
	PathParams PutAPITagsAddToOrderNumberPathParams
	Request    *PutAPITagsAddToOrderNumberRequestBody `request:"mediaType=application/json"`
}

type PutAPITagsAddToOrderNumberResponse struct {
	ContentType string
	StatusCode  int64
}
