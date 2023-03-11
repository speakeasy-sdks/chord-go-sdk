package operations

import (
	"net/http"
)

type PostAPISolidusImporterImportsRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

type PostAPISolidusImporterImportsRequest struct {
	Request *PostAPISolidusImporterImportsRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
}

type PostAPISolidusImporterImportsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
