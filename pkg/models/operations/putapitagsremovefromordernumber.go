package operations

type PutAPITagsRemoveFromOrderNumberPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type PutAPITagsRemoveFromOrderNumberRequestBody struct {
	Tags []string `json:"tags,omitempty"`
}

type PutAPITagsRemoveFromOrderNumberRequest struct {
	PathParams PutAPITagsRemoveFromOrderNumberPathParams
	Request    *PutAPITagsRemoveFromOrderNumberRequestBody `request:"mediaType=application/json"`
}

type PutAPITagsRemoveFromOrderNumberResponse struct {
	ContentType string
	StatusCode  int64
}
