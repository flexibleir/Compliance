// Code generated by go-swagger; DO NOT EDIT.

package compliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetIDHandlerFunc turns a function with the right signature into a get ID handler
type GetIDHandlerFunc func(GetIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIDHandlerFunc) Handle(params GetIDParams) middleware.Responder {
	return fn(params)
}

// GetIDHandler interface for that can handle valid get ID params
type GetIDHandler interface {
	Handle(GetIDParams) middleware.Responder
}

// NewGetID creates a new http.Handler for the get ID operation
func NewGetID(ctx *middleware.Context, handler GetIDHandler) *GetID {
	return &GetID{Context: ctx, Handler: handler}
}

/*GetID swagger:route GET /{id} compliance getId

GetID get ID API

*/
type GetID struct {
	Context *middleware.Context
	Handler GetIDHandler
}

func (o *GetID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
