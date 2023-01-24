package operations

type PutAPITagsIDRemoveFromOrderNumberPathParams struct {
	ID          int64  `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type PutAPITagsIDRemoveFromOrderNumberRequest struct {
	PathParams PutAPITagsIDRemoveFromOrderNumberPathParams
}

type PutAPITagsIDRemoveFromOrderNumberResponse struct {
	ContentType string
	StatusCode  int64
}
