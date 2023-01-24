package operations

type PutAPITagsAddToOrdersRequestBody struct {
	Orders []string `json:"orders,omitempty"`
	Tags   []string `json:"tags,omitempty"`
}

type PutAPITagsAddToOrdersRequest struct {
	Request *PutAPITagsAddToOrdersRequestBody `request:"mediaType=application/json"`
}

type PutAPITagsAddToOrdersResponse struct {
	ContentType string
	StatusCode  int64
}
