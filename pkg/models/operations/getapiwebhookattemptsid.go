package operations

type GetAPIWebhookAttemptsIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAPIWebhookAttemptsIDRequest struct {
	PathParams GetAPIWebhookAttemptsIDPathParams
}

type GetAPIWebhookAttemptsID200ApplicationJSON struct {
	EventName      *string `json:"event_name,omitempty"`
	HTTPStatusCode *int64  `json:"http_status_code,omitempty"`
	ID             *int64  `json:"id,omitempty"`
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	TargetURL      *string `json:"target_url,omitempty"`
	Token          *string `json:"token,omitempty"`
}

type GetAPIWebhookAttemptsIDResponse struct {
	ContentType                                     string
	StatusCode                                      int64
	GetAPIWebhookAttemptsID200ApplicationJSONObject *GetAPIWebhookAttemptsID200ApplicationJSON
}
