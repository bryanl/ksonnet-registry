// Code generated by go-swagger; DO NOT EDIT.

package blobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PullBlobHandlerFunc turns a function with the right signature into a pull blob handler
type PullBlobHandlerFunc func(PullBlobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PullBlobHandlerFunc) Handle(params PullBlobParams) middleware.Responder {
	return fn(params)
}

// PullBlobHandler interface for that can handle valid pull blob params
type PullBlobHandler interface {
	Handle(PullBlobParams) middleware.Responder
}

// NewPullBlob creates a new http.Handler for the pull blob operation
func NewPullBlob(ctx *middleware.Context, handler PullBlobHandler) *PullBlob {
	return &PullBlob{Context: ctx, Handler: handler}
}

/*PullBlob swagger:route GET /api/v1/packages/{namespace}/{package}/blobs/sha256/{digest} blobs pullBlob

Pull a package blob by digest

*/
type PullBlob struct {
	Context *middleware.Context
	Handler PullBlobHandler
}

func (o *PullBlob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPullBlobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
