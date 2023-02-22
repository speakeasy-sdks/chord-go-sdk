package operations

type UpdateSubscriptionTagsPathParams struct {
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscription_id"`
}

type UpdateSubscriptionTagsRequestBody struct {
	Tags []string `json:"tags,omitempty"`
}

type UpdateSubscriptionTagsRequest struct {
	PathParams UpdateSubscriptionTagsPathParams
	Request    *UpdateSubscriptionTagsRequestBody `request:"mediaType=application/json"`
}

type UpdateSubscriptionTagsResponse struct {
	ContentType string
	StatusCode  int
}
