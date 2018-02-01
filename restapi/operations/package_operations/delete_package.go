// Code generated by go-swagger; DO NOT EDIT.

package package_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeletePackageHandlerFunc turns a function with the right signature into a delete package handler
type DeletePackageHandlerFunc func(DeletePackageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePackageHandlerFunc) Handle(params DeletePackageParams) middleware.Responder {
	return fn(params)
}

// DeletePackageHandler interface for that can handle valid delete package params
type DeletePackageHandler interface {
	Handle(DeletePackageParams) middleware.Responder
}

// NewDeletePackage creates a new http.Handler for the delete package operation
func NewDeletePackage(ctx *middleware.Context, handler DeletePackageHandler) *DeletePackage {
	return &DeletePackage{Context: ctx, Handler: handler}
}

/*DeletePackage swagger:route DELETE /api/v1/packages/{namespace}/{package}/{release} package deletePackage

Delete a package release

*/
type DeletePackage struct {
	Context *middleware.Context
	Handler DeletePackageHandler
}

func (o *DeletePackage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeletePackageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
