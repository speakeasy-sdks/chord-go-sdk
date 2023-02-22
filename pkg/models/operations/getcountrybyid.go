package operations

type GetCountryByIDPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCountryByIDRequest struct {
	PathParams GetCountryByIDPathParams
}

type GetCountryByIDResponse struct {
	ContentType string
	StatusCode  int
}
