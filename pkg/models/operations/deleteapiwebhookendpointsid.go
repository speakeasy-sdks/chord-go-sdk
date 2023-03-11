package operations

import (
	"net/http"
)

type DeleteAPIWebhookEndpointsIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteAPIWebhookEndpointsIDRequest struct {
	PathParams DeleteAPIWebhookEndpointsIDPathParams
}

type DeleteAPIWebhookEndpointsIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
