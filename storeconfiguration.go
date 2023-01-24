package sdk

import (
	"context"
	"fmt"
	"net/http"
	"openapi/pkg/models/operations"
	"openapi/pkg/utils"
)

type StoreConfiguration struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewStoreConfiguration(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *StoreConfiguration {
	return &StoreConfiguration{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// GetAPIStoresStoreIDEnvironmentVariables - get Store Configuration
// Gets the Store Configuration
func (s *StoreConfiguration) GetAPIStoresStoreIDEnvironmentVariables(ctx context.Context, request operations.GetAPIStoresStoreIDEnvironmentVariablesRequest) (*operations.GetAPIStoresStoreIDEnvironmentVariablesResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/stores/{store_id}/environment_variables", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIStoresStoreIDEnvironmentVariablesResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetAPIStoresStoreIDEnvironmentVariables200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PatchAPIStoresStoreIDEnvironmentVariables - Update Store Configuration
// Update the Store Configuration
func (s *StoreConfiguration) PatchAPIStoresStoreIDEnvironmentVariables(ctx context.Context, request operations.PatchAPIStoresStoreIDEnvironmentVariablesRequest) (*operations.PatchAPIStoresStoreIDEnvironmentVariablesResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/stores/{store_id}/environment_variables", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchAPIStoresStoreIDEnvironmentVariablesResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PatchAPIStoresStoreIDEnvironmentVariables201ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PostAPIStoresStoreIDEnvironmentVariables - Update Store Configuration
// Update the Store Configuration
func (s *StoreConfiguration) PostAPIStoresStoreIDEnvironmentVariables(ctx context.Context, request operations.PostAPIStoresStoreIDEnvironmentVariablesRequest) (*operations.PostAPIStoresStoreIDEnvironmentVariablesResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/stores/{store_id}/environment_variables", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIStoresStoreIDEnvironmentVariablesResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PostAPIStoresStoreIDEnvironmentVariables201ApplicationJSONObject = out
		}
	}

	return res, nil
}
