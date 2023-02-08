package operations

type GetAPIStoresStoreIDEnvironmentVariablesPathParams struct {
	StoreID int64 `pathParam:"style=simple,explode=false,name=store_id"`
}

type GetAPIStoresStoreIDEnvironmentVariablesRequest struct {
	PathParams GetAPIStoresStoreIDEnvironmentVariablesPathParams
}

// GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSON
// The Environment Variables of the Store
type GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSON struct {
	Config1 *string `json:"CONFIG_1,omitempty"`
	Config2 *string `json:"CONFIG_2,omitempty"`
}

type GetAPIStoresStoreIDEnvironmentVariablesResponse struct {
	ContentType                                                     string
	StatusCode                                                      int64
	GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSONObject *GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSON
}
