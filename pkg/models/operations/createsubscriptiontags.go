package operations

type CreateSubscriptionTagsPathParams struct {
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscription_id"`
}

type CreateSubscriptionTagsRequestBody struct {
	Tags []string `json:"tags,omitempty"`
}

type CreateSubscriptionTagsRequest struct {
	PathParams CreateSubscriptionTagsPathParams
	Request    *CreateSubscriptionTagsRequestBody `request:"mediaType=application/json"`
}

type CreateSubscriptionTagsResponse struct {
	ContentType string
	StatusCode  int64
}
