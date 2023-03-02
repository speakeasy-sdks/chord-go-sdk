package sdk

import (
	"context"
	"fmt"
	"net/http"
	"openapi/v2/pkg/models/operations"
	"openapi/v2/pkg/utils"
	"strings"
)

type endpoint struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newEndpoint(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *endpoint {
	return &endpoint{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateEndpoint - create Endpoint
// if an endpoint with the same name exists, it is simply returned
func (s *endpoint) CreateEndpoint(ctx context.Context, request operations.CreateEndpointRequest) (*operations.CreateEndpointResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/webhook/endpoints"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
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

	res := &operations.CreateEndpointResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// DeleteAPIWebhookEndpointsID - Remove Endpoint
func (s *endpoint) DeleteAPIWebhookEndpointsID(ctx context.Context, request operations.DeleteAPIWebhookEndpointsIDRequest) (*operations.DeleteAPIWebhookEndpointsIDResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/webhook/endpoints/{id}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

	res := &operations.DeleteAPIWebhookEndpointsIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}
