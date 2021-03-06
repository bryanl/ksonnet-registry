// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bryanl/ksonnet-registry/restapi/operations/blobs"
	"github.com/bryanl/ksonnet-registry/restapi/operations/info"
	"github.com/bryanl/ksonnet-registry/restapi/operations/package_operations"
)

// NewKsonnetRegistryAPI creates a new KsonnetRegistry instance
func NewKsonnetRegistryAPI(spec *loads.Document) *KsonnetRegistryAPI {
	return &KsonnetRegistryAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		GzipProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("gzip producer has not yet been implemented")
		}),
		PackageOperationsCreatePackageHandler: package_operations.CreatePackageHandlerFunc(func(params package_operations.CreatePackageParams) middleware.Responder {
			return middleware.NotImplemented("operation PackageOperationsCreatePackage has not yet been implemented")
		}),
		PackageOperationsDeletePackageReleaseHandler: package_operations.DeletePackageReleaseHandlerFunc(func(params package_operations.DeletePackageReleaseParams) middleware.Responder {
			return middleware.NotImplemented("operation PackageOperationsDeletePackageRelease has not yet been implemented")
		}),
		InfoGetVersionHandler: info.GetVersionHandlerFunc(func(params info.GetVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation InfoGetVersion has not yet been implemented")
		}),
		PackageOperationsListPackagesHandler: package_operations.ListPackagesHandlerFunc(func(params package_operations.ListPackagesParams) middleware.Responder {
			return middleware.NotImplemented("operation PackageOperationsListPackages has not yet been implemented")
		}),
		BlobsPullBlobHandler: blobs.PullBlobHandlerFunc(func(params blobs.PullBlobParams) middleware.Responder {
			return middleware.NotImplemented("operation BlobsPullBlob has not yet been implemented")
		}),
		PullPackageHandler: PullPackageHandlerFunc(func(params PullPackageParams) middleware.Responder {
			return middleware.NotImplemented("operation PullPackage has not yet been implemented")
		}),
		PackageOperationsShowPackageReleaseHandler: package_operations.ShowPackageReleaseHandlerFunc(func(params package_operations.ShowPackageReleaseParams) middleware.Responder {
			return middleware.NotImplemented("operation PackageOperationsShowPackageRelease has not yet been implemented")
		}),
		PackageOperationsShowPackageReleasesHandler: package_operations.ShowPackageReleasesHandlerFunc(func(params package_operations.ShowPackageReleasesParams) middleware.Responder {
			return middleware.NotImplemented("operation PackageOperationsShowPackageReleases has not yet been implemented")
		}),
	}
}

/*KsonnetRegistryAPI Ksonnet Registry API documentation
 */
type KsonnetRegistryAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// GzipProducer registers a producer for a "application/x-gzip" mime type
	GzipProducer runtime.Producer

	// PackageOperationsCreatePackageHandler sets the operation handler for the create package operation
	PackageOperationsCreatePackageHandler package_operations.CreatePackageHandler
	// PackageOperationsDeletePackageReleaseHandler sets the operation handler for the delete package release operation
	PackageOperationsDeletePackageReleaseHandler package_operations.DeletePackageReleaseHandler
	// InfoGetVersionHandler sets the operation handler for the get version operation
	InfoGetVersionHandler info.GetVersionHandler
	// PackageOperationsListPackagesHandler sets the operation handler for the list packages operation
	PackageOperationsListPackagesHandler package_operations.ListPackagesHandler
	// BlobsPullBlobHandler sets the operation handler for the pull blob operation
	BlobsPullBlobHandler blobs.PullBlobHandler
	// PullPackageHandler sets the operation handler for the pull package operation
	PullPackageHandler PullPackageHandler
	// PackageOperationsShowPackageReleaseHandler sets the operation handler for the show package release operation
	PackageOperationsShowPackageReleaseHandler package_operations.ShowPackageReleaseHandler
	// PackageOperationsShowPackageReleasesHandler sets the operation handler for the show package releases operation
	PackageOperationsShowPackageReleasesHandler package_operations.ShowPackageReleasesHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *KsonnetRegistryAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *KsonnetRegistryAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *KsonnetRegistryAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *KsonnetRegistryAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *KsonnetRegistryAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *KsonnetRegistryAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *KsonnetRegistryAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the KsonnetRegistryAPI
func (o *KsonnetRegistryAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GzipProducer == nil {
		unregistered = append(unregistered, "GzipProducer")
	}

	if o.PackageOperationsCreatePackageHandler == nil {
		unregistered = append(unregistered, "package_operations.CreatePackageHandler")
	}

	if o.PackageOperationsDeletePackageReleaseHandler == nil {
		unregistered = append(unregistered, "package_operations.DeletePackageReleaseHandler")
	}

	if o.InfoGetVersionHandler == nil {
		unregistered = append(unregistered, "info.GetVersionHandler")
	}

	if o.PackageOperationsListPackagesHandler == nil {
		unregistered = append(unregistered, "package_operations.ListPackagesHandler")
	}

	if o.BlobsPullBlobHandler == nil {
		unregistered = append(unregistered, "blobs.PullBlobHandler")
	}

	if o.PullPackageHandler == nil {
		unregistered = append(unregistered, "PullPackageHandler")
	}

	if o.PackageOperationsShowPackageReleaseHandler == nil {
		unregistered = append(unregistered, "package_operations.ShowPackageReleaseHandler")
	}

	if o.PackageOperationsShowPackageReleasesHandler == nil {
		unregistered = append(unregistered, "package_operations.ShowPackageReleasesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *KsonnetRegistryAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *KsonnetRegistryAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *KsonnetRegistryAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *KsonnetRegistryAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *KsonnetRegistryAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "application/x-gzip":
			result["application/x-gzip"] = o.GzipProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *KsonnetRegistryAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the ksonnet registry API
func (o *KsonnetRegistryAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *KsonnetRegistryAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1/packages/{namespace}/{package}"] = package_operations.NewCreatePackage(o.context, o.PackageOperationsCreatePackageHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/v1/packages/{namespace}/{package}/{release}"] = package_operations.NewDeletePackageRelease(o.context, o.PackageOperationsDeletePackageReleaseHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/version"] = info.NewGetVersion(o.context, o.InfoGetVersionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/packages"] = package_operations.NewListPackages(o.context, o.PackageOperationsListPackagesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/packages/{namespace}/{package}/blobs/sha256/{digest}"] = blobs.NewPullBlob(o.context, o.BlobsPullBlobHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/packages/{namespace}/{package}/{release}/pull"] = NewPullPackage(o.context, o.PullPackageHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/packages/{namespace}/{package}/{release}"] = package_operations.NewShowPackageRelease(o.context, o.PackageOperationsShowPackageReleaseHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/packages/{namespace}/{package}"] = package_operations.NewShowPackageReleases(o.context, o.PackageOperationsShowPackageReleasesHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *KsonnetRegistryAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *KsonnetRegistryAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
