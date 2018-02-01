package apiadapter

import (
	"github.com/bryanl/ksonnet-registry/models"
	"github.com/bryanl/ksonnet-registry/restapi/operations/info"
	"github.com/go-openapi/runtime/middleware"
)

// GetVersion gets the service version.
func GetVersion(params info.GetVersionParams) middleware.Responder {
	payload := &models.Version{
		KsonnetRegistryAPI: "0.1.0",
	}

	resp := info.NewGetVersionOK().WithPayload(payload)

	return resp
}
