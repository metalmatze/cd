// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CurrentDeploymentHandlerFunc turns a function with the right signature into a current deployment handler
type CurrentDeploymentHandlerFunc func(CurrentDeploymentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CurrentDeploymentHandlerFunc) Handle(params CurrentDeploymentParams) middleware.Responder {
	return fn(params)
}

// CurrentDeploymentHandler interface for that can handle valid current deployment params
type CurrentDeploymentHandler interface {
	Handle(CurrentDeploymentParams) middleware.Responder
}

// NewCurrentDeployment creates a new http.Handler for the current deployment operation
func NewCurrentDeployment(ctx *middleware.Context, handler CurrentDeploymentHandler) *CurrentDeployment {
	return &CurrentDeployment{Context: ctx, Handler: handler}
}

/*CurrentDeployment swagger:route GET /deployments/current deployments currentDeployment

Returns the currently active deployment

*/
type CurrentDeployment struct {
	Context *middleware.Context
	Handler CurrentDeploymentHandler
}

func (o *CurrentDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCurrentDeploymentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
