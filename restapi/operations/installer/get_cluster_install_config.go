// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClusterInstallConfigHandlerFunc turns a function with the right signature into a get cluster install config handler
type GetClusterInstallConfigHandlerFunc func(GetClusterInstallConfigParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetClusterInstallConfigHandlerFunc) Handle(params GetClusterInstallConfigParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetClusterInstallConfigHandler interface for that can handle valid get cluster install config params
type GetClusterInstallConfigHandler interface {
	Handle(GetClusterInstallConfigParams, interface{}) middleware.Responder
}

// NewGetClusterInstallConfig creates a new http.Handler for the get cluster install config operation
func NewGetClusterInstallConfig(ctx *middleware.Context, handler GetClusterInstallConfigHandler) *GetClusterInstallConfig {
	return &GetClusterInstallConfig{Context: ctx, Handler: handler}
}

/*GetClusterInstallConfig swagger:route GET /clusters/{cluster_id}/install-config installer getClusterInstallConfig

Get the cluster install config yaml

*/
type GetClusterInstallConfig struct {
	Context *middleware.Context
	Handler GetClusterInstallConfigHandler
}

func (o *GetClusterInstallConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetClusterInstallConfigParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}