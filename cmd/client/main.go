package main

import (
	"context"
	"net/http"

	"github.com/bryanl/ksonnet-registry/client"
	"github.com/bryanl/ksonnet-registry/client/package_operations"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func main() {
	config := &client.TransportConfig{
		Host:     "localhost:9000",
		BasePath: client.DefaultBasePath,
		Schemes:  client.DefaultSchemes,
	}

	c := client.NewHTTPClientWithConfig(nil, config)

	ctx := context.Background()

	params := &package_operations.ShowPackageReleasesParams{
		Context:    ctx,
		HTTPClient: http.DefaultClient,
		Namespace:  "ns",
		Package:    "node",
	}

	releases, err := c.PackageOperations.ShowPackageReleases(params)
	if err != nil {
		logrus.WithError(err).Fatal("show package releases")
	}

	spew.Dump(releases.Payload)
}
