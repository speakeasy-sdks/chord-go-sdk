package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-go-sdk/pkg/utils"
	"net/http"
	"strings"
)

type StockRequest struct {
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

func NewStockRequest(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *StockRequest {
	return &StockRequest{
		_defaultClient:  defaultClient,
		_securityClient: securityClient,
		_serverURL:      serverURL,
		_language:       language,
		_sdkVersion:     sdkVersion,
		_genVersion:     genVersion,
	}
}

// CreateStockRequest - Create Stock Request
func (s *StockRequest) CreateStockRequest(ctx context.Context, request operations.CreateStockRequestRequest) (*operations.CreateStockRequestResponse, error) {
	baseURL := s._serverURL
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

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateStockRequestResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
	case httpRes.StatusCode == 422:
	}

	return res, nil
}
