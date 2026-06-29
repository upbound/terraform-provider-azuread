package xpprovider

import (
	"context"

	sdkclient "github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/provider"
)

type AzureADClientBuilder clients.ClientBuilder

func GetProviderSchema(_ context.Context) (*schema.Provider, error) {
	return provider.AzureADProvider(), nil
}

func (b *AzureADClientBuilder) GetClient(ctx context.Context) (*clients.Client, error) {
	return (*clients.ClientBuilder)(b).Build(ctx)
}

// RegisterResponseMiddleware appends mw to every Microsoft Graph SDK client held
// by meta (the value returned by schema.Provider.Meta()). It returns false when
// meta is not a *clients.Client.
func RegisterResponseMiddleware(meta any, mw sdkclient.ResponseMiddleware) bool {
	c, ok := meta.(*clients.Client)
	if !ok {
		return false
	}
	c.AppendResponseMiddleware(mw)
	return true
}
