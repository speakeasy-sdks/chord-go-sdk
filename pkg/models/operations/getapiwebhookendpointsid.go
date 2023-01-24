package operations

type GetAPIWebhookEndpointsIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAPIWebhookEndpointsIDQueryParams struct {
	AttemptsCount   *int64 `queryParam:"style=form,explode=true,name=attempts_count"`
	IncludeAttempts *bool  `queryParam:"style=form,explode=true,name=include_attempts"`
}

type GetAPIWebhookEndpointsID200ApplicationJSON struct {
	Events    []string `json:"events,omitempty"`
	Filters   []string `json:"filters,omitempty"`
	ID        *int64   `json:"id,omitempty"`
	TargetURL *string  `json:"target_url,omitempty"`
	Token     *string  `json:"token,omitempty"`
}

type GetAPIWebhookEndpointsIDRequest struct {
	PathParams  GetAPIWebhookEndpointsIDPathParams
	QueryParams GetAPIWebhookEndpointsIDQueryParams
}

type GetAPIWebhookEndpointsIDResponse struct {
	ContentType                                      string
	StatusCode                                       int64
	GetAPIWebhookEndpointsID200ApplicationJSONObject *GetAPIWebhookEndpointsID200ApplicationJSON
}
