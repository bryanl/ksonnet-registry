package restapi

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	graceful "github.com/tylerb/graceful"

	"github.com/bryanl/ksonnet-registry/apiadapter"
	"github.com/bryanl/ksonnet-registry/models"
	"github.com/bryanl/ksonnet-registry/restapi/operations"
	"github.com/bryanl/ksonnet-registry/restapi/operations/blobs"
	"github.com/bryanl/ksonnet-registry/restapi/operations/info"
	"github.com/bryanl/ksonnet-registry/restapi/operations/package_operations"
	"github.com/bryanl/ksonnet-registry/store"
)

var (
	s store.Store
)

func init() {
	var err error
	s, err = store.NewTempStore()
	if err != nil {
		logrus.WithError(err).Fatal("unable to initialize store")
	}
}

//go:generate swagger generate server --target .. --name ksonnet-registry --spec ../swagger.yml

func configureFlags(api *operations.KsonnetRegistryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KsonnetRegistryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = logrus.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GzipProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		switch t := data.(type) {
		default:
			return errors.NotImplemented(fmt.Sprintf("not sure what to do with file of type %T", t))
		case models.PullBlobOKBody:
			if _, err := io.Copy(w, t.Data); err != nil {
				return err
			}

			return nil
		}
	})

	api.PackageOperationsCreatePackageHandler = package_operations.CreatePackageHandlerFunc(func(params package_operations.CreatePackageParams) middleware.Responder {
		return apiadapter.CreatePackage(s, params)
	})
	api.PackageOperationsDeletePackageReleaseHandler = package_operations.DeletePackageReleaseHandlerFunc(func(params package_operations.DeletePackageReleaseParams) middleware.Responder {
		return apiadapter.DeletePackageRelease(s, params)
	})
	api.InfoGetVersionHandler = info.GetVersionHandlerFunc(func(params info.GetVersionParams) middleware.Responder {
		return middleware.NotImplemented("operation info.GetVersion has not yet been implemented")
	})
	api.PackageOperationsListPackagesHandler = package_operations.ListPackagesHandlerFunc(func(params package_operations.ListPackagesParams) middleware.Responder {
		return middleware.NotImplemented("operation package_operations.ListPackages has not yet been implemented")
	})
	api.BlobsPullBlobHandler = blobs.PullBlobHandlerFunc(func(params blobs.PullBlobParams) middleware.Responder {
		return apiadapter.PullPackage(s, params)
	})
	api.PullPackageHandler = operations.PullPackageHandlerFunc(func(params operations.PullPackageParams) middleware.Responder {
		return middleware.NotImplemented("operation .PullPackage has not yet been implemented")
	})
	api.PackageOperationsShowPackageReleasesHandler = package_operations.ShowPackageReleasesHandlerFunc(func(params package_operations.ShowPackageReleasesParams) middleware.Responder {
		return apiadapter.ShowPackage(s, params)
	})
	api.PackageOperationsShowPackageReleaseHandler = package_operations.ShowPackageReleaseHandlerFunc(func(params package_operations.ShowPackageReleaseParams) middleware.Responder {
		return middleware.NotImplemented("show package release handler not implemented")
	})

	api.ServerShutdown = func() {
		s.Close()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
