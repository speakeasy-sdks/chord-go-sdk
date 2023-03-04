package operations

import (
	"net/http"
)

type PostAPIStoresStoreIDEnvironmentVariablesPathParams struct {
	StoreID int64 `pathParam:"style=simple,explode=false,name=store_id"`
}

type PostAPIStoresStoreIDEnvironmentVariablesRequestBody struct {
	Settings []interface{} `json:"settings"`
}

type PostAPIStoresStoreIDEnvironmentVariablesRequest struct {
	PathParams PostAPIStoresStoreIDEnvironmentVariablesPathParams
	Request    *PostAPIStoresStoreIDEnvironmentVariablesRequestBody `request:"mediaType=application/json"`
}

// PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
// The updated Environment Variables of the Store
type PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSON struct {
	Config1 *string `json:"CONFIG_1,omitempty"`
	Config2 *string `json:"CONFIG_2,omitempty"`
}

type PostAPIStoresStoreIDEnvironmentVariablesResponse struct {
	ContentType                                                      string
	StatusCode                                                       int
	RawResponse                                                      *http.Response
	PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSONObject *PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
}
