package operations

type CreateStockRequestRequestBodyStockRequest struct {
	Email string `json:"email"`
	Sku   string `json:"sku"`
}

type CreateStockRequestRequestBody struct {
	StockRequest CreateStockRequestRequestBodyStockRequest `json:"stock_request"`
}

type CreateStockRequestRequest struct {
	Request *CreateStockRequestRequestBody `request:"mediaType=application/json"`
}

type CreateStockRequestResponse struct {
	ContentType string
	StatusCode  int
}
