package operations

type CreateTagRequestBodyTag struct {
	Name string `json:"name"`
}

type CreateTagRequestBody struct {
	Tag CreateTagRequestBodyTag `json:"tag"`
}

type CreateTagRequest struct {
	Request *CreateTagRequestBody `request:"mediaType=application/json"`
}

type CreateTagResponse struct {
	ContentType string
	StatusCode  int
}
