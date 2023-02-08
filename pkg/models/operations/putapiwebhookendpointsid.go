package operations

type PutAPIWebhookEndpointsIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type PutAPIWebhookEndpointsIDRequestBodyEndpoint struct {
	TargetURL string  `json:"target_url"`
	Token     *string `json:"token,omitempty"`
}

type PutAPIWebhookEndpointsIDRequestBody struct {
	Endpoint PutAPIWebhookEndpointsIDRequestBodyEndpoint `json:"endpoint"`
}

type PutAPIWebhookEndpointsIDRequest struct {
	PathParams PutAPIWebhookEndpointsIDPathParams
	Request    *PutAPIWebhookEndpointsIDRequestBody `request:"mediaType=application/json"`
}

type PutAPIWebhookEndpointsID200ApplicationJSON struct {
	Events    []string `json:"events,omitempty"`
	ID        *int64   `json:"id,omitempty"`
	TargetURL *string  `json:"target_url,omitempty"`
	Token     *string  `json:"token,omitempty"`
}

type PutAPIWebhookEndpointsIDResponse struct {
	ContentType                                      string
	StatusCode                                       int64
	PutAPIWebhookEndpointsID200ApplicationJSONObject *PutAPIWebhookEndpointsID200ApplicationJSON
}
