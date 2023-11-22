package xpprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/provider"
)

func GetProviderSchema(_ context.Context) (*schema.Provider, error) {
	return provider.AzureADProvider(), nil
}
