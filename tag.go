package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/chord-go-sdk/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/chord-go-sdk/v2/pkg/utils"
	"net/http"
	"strings"
)

type tag struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newTag(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *tag {
	return &tag{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateTag - create Tag
// if a tag with the same name exists, it is simply returned
func (s *tag) CreateTag(ctx context.Context, request operations.CreateTagRequest) (*operations.CreateTagResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/api/tags"

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

	res := &operations.CreateTagResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
