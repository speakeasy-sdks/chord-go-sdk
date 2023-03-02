package sdk

import (
	"context"
	"fmt"
	"net/http"
	"openapi/v2/pkg/models/operations"
	"openapi/v2/pkg/utils"
	"strings"
)

type referralIdentifier struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newReferralIdentifier(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *referralIdentifier {
	return &referralIdentifier{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// FindOrCreateReferralIdentifier - Referral identifier
func (s *referralIdentifier) FindOrCreateReferralIdentifier(ctx context.Context, request operations.FindOrCreateReferralIdentifierRequest) (*operations.FindOrCreateReferralIdentifierResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/referral_identifiers"

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

	res := &operations.FindOrCreateReferralIdentifierResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	case httpRes.StatusCode == 201:
	}

	return res, nil
}
