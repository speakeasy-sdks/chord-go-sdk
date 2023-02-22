package operations

type GetStatesByIDPathParams struct {
	CountryID string `pathParam:"style=simple,explode=false,name=country_id"`
}

type GetStatesByIDRequest struct {
	PathParams GetStatesByIDPathParams
}

type GetStatesByIDResponse struct {
	ContentType string
	StatusCode  int
}
