package operations

type GetAPIStoresStoreIDEnvironmentVariablesPathParams struct {
	StoreID int64 `pathParam:"style=simple,explode=false,name=store_id"`
}

// GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSON
// The Environment Variables of the Store
type GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSON struct {
	CONFIG1 *string `json:"CONFIG_1,omitempty"`
	CONFIG2 *string `json:"CONFIG_2,omitempty"`
}

type GetAPIStoresStoreIDEnvironmentVariablesRequest struct {
	PathParams GetAPIStoresStoreIDEnvironmentVariablesPathParams
}

type GetAPIStoresStoreIDEnvironmentVariablesResponse struct {
	ContentType                                                     string
	StatusCode                                                      int64
	GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSONObject *GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSON
}
