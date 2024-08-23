# bluestonepim-go-sdk

[![Build Status](https://github.com/labd/bluestonepim-go-sdk/workflows/Go%20Tests/badge.svg)](https://github.com/labd/bluestonepim-go-sdk/workflows/)
[![codecov](https://codecov.io/gh/LabD/bluestonepim-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/LabD/bluestonepim-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/labd/bluestonepim-go-sdk)](https://goreportcard.com/report/github.com/labd/bluestonepim-go-sdk)
[![GoDoc](https://godoc.org/github.com/labd/bluestonepim-go-sdk?status.svg)](https://godoc.org/github.com/labd/bluestonepim-go-sdk)

The Bluestone PIM Go SDK is automatically generated based on the official [API specifications](https://docs.api.bluestonepim.com/docs/openapi-documentation)
of Bluestone PIM. It should therefore be nearly feature complete.

The SDK was initially created for enabling the creation of the
[Terraform Provider for Bluestone PIM](https://github.com/labd/terraform-provider-bluestonepim)
That provider enables you to use infrastructure-as-code principles with Bluestone PIM.

Note that since this SDK is automatically generated we cannot guarantee backwards
compatibility between releases. Please pin the dependency correctly and be aware
of potential changes when updating

## Using the SDK

```go
package main

import (
   "context"
   "errors"
   "fmt"
   "github.com/labd/bluestonepim-go-sdk/pim"
   "log"
   "math/rand"
   "time"

   "golang.org/x/oauth2/clientcredentials"
)

func main() {

   // Create the new client. When an empty value is passed it will use the CTP_*
   // environment variables to get the value. The HTTPClient arg is optional,
   // and when empty will automatically be created using the env values.
   oauth2Config := &clientcredentials.Config{
      ClientID:     "<your-client-id>",
      ClientSecret: "<your-client-secret",
      TokenURL:     "https://idp.bluestonepim.com/op/token",
   }
   
   ctx := context.Background()

   httpClient := oauth2Config.Client(ctx)

   client, err := pim.NewClientWithResponses(
      fmt.Sprintf("%s/pim", "https://api.bluestonepim.com"),
      pim.WithHTTPClient(httpClient),
   )
   if err != nil {
      log.Fatal(err)
   }
   
   // Create a new category
   response, err := client.CreateCategoryWithResponse(ctx,
      &pim.CreateCategoryParams{
         Validation: "NAME",
      },
      pim.CreateCategoryJSONRequestBody{
         Name:     "my category name",
         Number:   "my-category-key",
      },
   )
    if err != nil {
        log.Fatal(err)
    }
}

```

## Generating code

To re-generate the API take the following steps:
 - Install `task` from https://taskfile.dev/installation/
 - `go mod download` to fetch all the required tools
 - `task generate` to generate the code
