///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAPIHandlerFunc turns a function with the right signature into a get API handler
type GetAPIHandlerFunc func(GetAPIParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAPIHandlerFunc) Handle(params GetAPIParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAPIHandler interface for that can handle valid get API params
type GetAPIHandler interface {
	Handle(GetAPIParams, interface{}) middleware.Responder
}

// NewGetAPI creates a new http.Handler for the get API operation
func NewGetAPI(ctx *middleware.Context, handler GetAPIHandler) *GetAPI {
	return &GetAPI{Context: ctx, Handler: handler}
}

/*GetAPI swagger:route GET /{api} endpoint getApi

Find API by name

get an API by name

*/
type GetAPI struct {
	Context *middleware.Context
	Handler GetAPIHandler
}

func (o *GetAPI) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAPIParams()

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