package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-go-sdk/v2/pkg/utils"
	"net/http"
	"strings"
)

type attempts struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newAttempts(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *attempts {
	return &attempts{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// FindWebhookAttemps - get all Webhook Attemps
// get all webhook attempts
func (s *attempts) FindWebhookAttemps(ctx context.Context, request operations.FindWebhookAttempsRequest) (*operations.FindWebhookAttempsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/webhook/attempts"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

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

	res := &operations.FindWebhookAttempsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAPIWebhookAttemptsID - retrieve Attempt
func (s *attempts) GetAPIWebhookAttemptsID(ctx context.Context, request operations.GetAPIWebhookAttemptsIDRequest) (*operations.GetAPIWebhookAttemptsIDResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/api/webhook/attempts/{id}", request.PathParams)

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

	res := &operations.GetAPIWebhookAttemptsIDResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetAPIWebhookAttemptsID200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetAPIWebhookAttemptsID200ApplicationJSONObject = out
		}
	}

	return res, nil
}
