package operations

type FindWebhookEndpointsQueryParams struct {
	AttemptsCount   *int64  `queryParam:"style=form,explode=true,name=attempts_count"`
	IncludeAttempts *bool   `queryParam:"style=form,explode=true,name=include_attempts"`
	Q               *string `queryParam:"style=form,explode=true,name=q"`
}

type FindWebhookEndpointsRequest struct {
	QueryParams FindWebhookEndpointsQueryParams
}

type FindWebhookEndpointsResponse struct {
	ContentType string
	StatusCode  int64
}
