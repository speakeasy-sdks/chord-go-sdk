package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-go-sdk/v2/pkg/utils"
	"net/http"
	"strings"
)

type reminders struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newReminders(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *reminders {
	return &reminders{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// FindReminders - get all Subscription Reminders
// get all subscription reminders
func (s *reminders) FindReminders(ctx context.Context, request operations.FindRemindersRequest) (*operations.FindRemindersResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/subscriptions/reminders"

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

	res := &operations.FindRemindersResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
