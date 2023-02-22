package operations

type GetSubscriptionsQueryParams struct {
	Q *string `queryParam:"style=form,explode=true,name=q"`
}

type GetSubscriptionsRequest struct {
	QueryParams GetSubscriptionsQueryParams
}

type GetSubscriptionsResponse struct {
	ContentType string
	StatusCode  int
}
