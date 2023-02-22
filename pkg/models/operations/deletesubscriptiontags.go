package operations

type DeleteSubscriptionTagsPathParams struct {
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscription_id"`
}

type DeleteSubscriptionTagsRequest struct {
	PathParams DeleteSubscriptionTagsPathParams
}

type DeleteSubscriptionTagsResponse struct {
	ContentType string
	StatusCode  int
}
