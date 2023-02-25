# Chord Commerce Go SDK

Chord’s headless e-commerce and data platform empower you to build bespoke, best-in-class e-commerce experiences. This is a Go SDK for the [Chord API](https://chord.stoplight.io/docs/chord-oms/f593af8ec51a9-chord-api)

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/chord-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
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

## Authentication

The API provides two means of authentication: one is through your user's API key, while the other is through an order's guest token.

### API key
You can use your API key to access all resources in the API. The API key must be passed in the Authorization header in the following form:

`Authorization: Bearer API_KEY`

As an admin, you can find your API token in the admin section under Users > Your e-email > API Access (at admin/users/<user_id>/edit)

<!-- Start API Key Example Usage -->
Example:

```bash
curl --header "Authorization: Bearer 1a6a9936ad150a2ee345c65331da7a3ccc2de" http://plant-staging.assembly-api.com/api/stores
```
<!-- End API Key Example Usage -->

### Order token
For allowing guests to manage their cart and place their order, you can use the order's guest token. This token is contained in the guest_token property of the order, and it allows you to perform certain checkout-related operations on the order such as managing line items, completing the checkout flow etc.

The order token must be passed in the X-Spree-Order-Token header in the following form:

`X-Spree-Order-Token: ORDER_TOKEN`

If you are already providing an API key, you don't need to also provide the order token (although you may do so). More information on authentication can be found [here](https://chord.stoplight.io/docs/chord-oms/ZG9jOjEwODE5NTQ-authentication)

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

### StoreConfiguration

* `GetAPIStoresStoreIDEnvironmentVariables` - get Store Configuration
* `PatchAPIStoresStoreIDEnvironmentVariables` - Update Store Configuration
* `PostAPIStoresStoreIDEnvironmentVariables` - Update Store Configuration

### SubscriptionTags

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
