package operations

type PatchAPIStoresStoreIDEnvironmentVariablesPathParams struct {
	StoreID int64 `pathParam:"style=simple,explode=false,name=store_id"`
}

type PatchAPIStoresStoreIDEnvironmentVariablesRequestBody struct {
	Settings []interface{} `json:"settings"`
}

// PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
// The updated Environment Variables of the Store
type PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSON struct {
	CONFIG1 *string `json:"CONFIG_1,omitempty"`
	CONFIG2 *string `json:"CONFIG_2,omitempty"`
}

type PatchAPIStoresStoreIDEnvironmentVariablesRequest struct {
	PathParams PatchAPIStoresStoreIDEnvironmentVariablesPathParams
	Request    *PatchAPIStoresStoreIDEnvironmentVariablesRequestBody `request:"mediaType=application/json"`
}

type PatchAPIStoresStoreIDEnvironmentVariablesResponse struct {
	ContentType                                                       string
	StatusCode                                                        int64
	PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSONObject *PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
}
