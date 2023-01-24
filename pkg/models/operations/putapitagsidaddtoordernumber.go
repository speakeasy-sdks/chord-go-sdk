package operations

type PutAPITagsIDAddToOrderNumberPathParams struct {
	ID          int64  `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type PutAPITagsIDAddToOrderNumberRequest struct {
	PathParams PutAPITagsIDAddToOrderNumberPathParams
}

type PutAPITagsIDAddToOrderNumberResponse struct {
	ContentType string
	StatusCode  int64
}
