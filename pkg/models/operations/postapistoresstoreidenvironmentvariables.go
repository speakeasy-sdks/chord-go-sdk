package operations

type PostAPIStoresStoreIDEnvironmentVariablesPathParams struct {
	StoreID int64 `pathParam:"style=simple,explode=false,name=store_id"`
}

type PostAPIStoresStoreIDEnvironmentVariablesRequestBody struct {
	Settings []interface{} `json:"settings"`
}

// PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
// The updated Environment Variables of the Store
type PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSON struct {
	CONFIG1 *string `json:"CONFIG_1,omitempty"`
	CONFIG2 *string `json:"CONFIG_2,omitempty"`
}

type PostAPIStoresStoreIDEnvironmentVariablesRequest struct {
	PathParams PostAPIStoresStoreIDEnvironmentVariablesPathParams
	Request    *PostAPIStoresStoreIDEnvironmentVariablesRequestBody `request:"mediaType=application/json"`
}

type PostAPIStoresStoreIDEnvironmentVariablesResponse struct {
	ContentType                                                      string
	StatusCode                                                       int64
	PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSONObject *PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
}
