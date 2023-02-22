package operations

type FindRemindersQueryParams struct {
	Q *string `queryParam:"style=form,explode=true,name=q"`
}

type FindRemindersRequest struct {
	QueryParams FindRemindersQueryParams
}

type FindRemindersResponse struct {
	ContentType string
	StatusCode  int
}
