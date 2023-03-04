package operations

import (
	"net/http"
)

type PatchAPIStoresStoreIDEnvironmentVariablesPathParams struct {
	StoreID int64 `pathParam:"style=simple,explode=false,name=store_id"`
}

type PatchAPIStoresStoreIDEnvironmentVariablesRequestBody struct {
	Settings []interface{} `json:"settings"`
}

type PatchAPIStoresStoreIDEnvironmentVariablesRequest struct {
	PathParams PatchAPIStoresStoreIDEnvironmentVariablesPathParams
	Request    *PatchAPIStoresStoreIDEnvironmentVariablesRequestBody `request:"mediaType=application/json"`
}

// PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
// The updated Environment Variables of the Store
type PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSON struct {
	Config1 *string `json:"CONFIG_1,omitempty"`
	Config2 *string `json:"CONFIG_2,omitempty"`
}

type PatchAPIStoresStoreIDEnvironmentVariablesResponse struct {
	ContentType                                                       string
	StatusCode                                                        int
	RawResponse                                                       *http.Response
	PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSONObject *PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
}
