# Chord Commerce Go SDK



<!-- Start SDK Installation -->
## SDK Installation

```bash
go get openapi
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "openapi"
    "openapi/pkg/models/shared"
    "openapi/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.FindWebhookAttempsRequest{
        QueryParams: operations.FindWebhookAttempsQueryParams{
            Q: "sit",
        },
    }
    
    res, err := s.Attempts.FindWebhookAttemps(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations

### Attempts

* `FindWebhookAttemps` - get all Webhook Attemps
* `GetAPIWebhookAttemptsID` - retrieve Attempt

### Country

* `GetCountryByID` - Get the country by id
* `ListCountries` - Get all countries

### Endpoint

* `CreateEndpoint` - create Endpoint
* `DeleteAPIWebhookEndpointsID` - Remove Endpoint

### Endpoints

* `FindWebhookEndpoints` - get all Webhook Endpoints
* `GetAPIWebhookEndpointsID` - retrieve Endpoint
* `PutAPIWebhookEndpointsID` - Update Endpoint

### Imports

* `PostAPISolidusImporterImports` - Create Import

### ReferralIdentifier

* `FindOrCreateReferralIdentifier` - Referral identifier

### Reminders

* `FindReminders` - get all Subscription Reminders

### Roles

* `CreateRole` - create Role
* `FindRoles` - get all Roles
* `GetAPIRolesID` - retrieve Role
* `PutAPIRolesID` - Update Role
* `PutAPIRolesIDAddUserID` - Add User to Role
* `PutAPIRolesIDRemoveUserID` - Remove User from Role

### State

* `GetStatesByID` - Get the states for one country retrieved by id
* `ListStates` - Get all states

### StockRequest

* `CreateStockRequest` - Create Stock Request

### Store Configuration

* `GetAPIStoresStoreIDEnvironmentVariables` - get Store Configuration
* `PatchAPIStoresStoreIDEnvironmentVariables` - Update Store Configuration
* `PostAPIStoresStoreIDEnvironmentVariables` - Update Store Configuration

### Subscription Tags

* `CreateSubscriptionTags` - add tags to subscription
* `DeleteSubscriptionTags` - remove tags from the subscription
* `GetSubscription` - Retrieve a subscription
* `GetSubscriptions` - Retrieves all subscriptions
* `UpdateSubscriptionTags` - set the tags on the subscription

### Tag

* `CreateTag` - create Tag

### Tags

* `FindTags` - get all Tags
* `GetAPITagsID` - retrieve Tag
* `PutAPITagsAddToOrderNumber` - Add Tag(s) to Order
* `PutAPITagsAddToOrders` - Add Tag(s) to Order(s)
* `PutAPITagsRemoveFromOrderNumber` - Remove Tag(s) from order
* `PutAPITagsID` - Update Tag
* `PutAPITagsIDRemoveFromOrderNumber` - Remove Tag(s) from order
* `PutAPITagsIDAddToOrderNumber` - Add Tag to Order

### WalletPaymentSource

* `ListWalletPaymentSource` - Get all user payment sources

<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
