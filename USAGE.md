<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/chord-go-sdk"
    "github.com/speakeasy-sdks/chord-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/chord-go-sdk/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            },
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.FindWebhookAttempsRequest{
        QueryParams: operations.FindWebhookAttempsQueryParams{
            Q: "unde",
        },
    }

    ctx := context.Background()
    res, err := s.Attempts.FindWebhookAttemps(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->