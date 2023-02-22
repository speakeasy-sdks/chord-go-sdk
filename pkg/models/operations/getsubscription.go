package operations

type GetSubscriptionPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetSubscriptionRequest struct {
	PathParams GetSubscriptionPathParams
}

type GetSubscriptionResponse struct {
	ContentType string
	StatusCode  int
}
