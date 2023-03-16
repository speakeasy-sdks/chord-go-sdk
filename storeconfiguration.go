package sdk

import (
	"context"
	"fmt"
	"net/http"
	"openapi/v2/pkg/models/operations"
	"openapi/v2/pkg/utils"
)

type storeConfiguration struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newStoreConfiguration(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *storeConfiguration {
	return &storeConfiguration{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetAPIStoresStoreIDEnvironmentVariables - get Store Configuration
// Gets the Store Configuration
func (s *storeConfiguration) GetAPIStoresStoreIDEnvironmentVariables(ctx context.Context, request operations.GetAPIStoresStoreIDEnvironmentVariablesRequest) (*operations.GetAPIStoresStoreIDEnvironmentVariablesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/stores/{store_id}/environment_variables", request.PathParams, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIStoresStoreIDEnvironmentVariablesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
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
func (s *storeConfiguration) PatchAPIStoresStoreIDEnvironmentVariables(ctx context.Context, request operations.PatchAPIStoresStoreIDEnvironmentVariablesRequest) (*operations.PatchAPIStoresStoreIDEnvironmentVariablesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/stores/{store_id}/environment_variables", request.PathParams, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchAPIStoresStoreIDEnvironmentVariablesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
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
func (s *storeConfiguration) PostAPIStoresStoreIDEnvironmentVariables(ctx context.Context, request operations.PostAPIStoresStoreIDEnvironmentVariablesRequest) (*operations.PostAPIStoresStoreIDEnvironmentVariablesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/stores/{store_id}/environment_variables", request.PathParams, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostAPIStoresStoreIDEnvironmentVariablesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
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
