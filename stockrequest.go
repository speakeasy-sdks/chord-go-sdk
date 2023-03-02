package sdk

import (
	"context"
	"fmt"
	"net/http"
	"openapi/v2/pkg/models/operations"
	"openapi/v2/pkg/utils"
	"strings"
)

type stockRequest struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newStockRequest(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *stockRequest {
	return &stockRequest{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateStockRequest - Create Stock Request
func (s *stockRequest) CreateStockRequest(ctx context.Context, request operations.CreateStockRequestRequest) (*operations.CreateStockRequestResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/stock_requests"

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

	res := &operations.CreateStockRequestResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
	case httpRes.StatusCode == 422:
	}

	return res, nil
}
