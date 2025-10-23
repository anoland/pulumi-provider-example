package examples

import (
	"github.com/anoland/pulumi-provider-example/provider"
	"github.com/pulumi/providertest/providers"
	goprovider "github.com/pulumi/pulumi-go-provider"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

var providerFactory = func(_ providers.PulumiTest) (pulumirpc.ResourceProviderServer, error) {
	return goprovider.RawServer("provider-boilerplate", "1.0.0", provider.Provider())(nil)
}
