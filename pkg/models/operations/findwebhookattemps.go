package operations

type FindWebhookAttempsQueryParams struct {
	Q *string `queryParam:"style=form,explode=true,name=q"`
}

type FindWebhookAttempsRequest struct {
	QueryParams FindWebhookAttempsQueryParams
}

type FindWebhookAttempsResponse struct {
	ContentType string
	StatusCode  int64
}
