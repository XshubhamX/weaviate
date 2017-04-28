/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package acl_entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateACLEntriesDeleteHandlerFunc turns a function with the right signature into a weaviate acl entries delete handler
type WeaviateACLEntriesDeleteHandlerFunc func(WeaviateACLEntriesDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateACLEntriesDeleteHandlerFunc) Handle(params WeaviateACLEntriesDeleteParams) middleware.Responder {
	return fn(params)
}

// WeaviateACLEntriesDeleteHandler interface for that can handle valid weaviate acl entries delete params
type WeaviateACLEntriesDeleteHandler interface {
	Handle(WeaviateACLEntriesDeleteParams) middleware.Responder
}

// NewWeaviateACLEntriesDelete creates a new http.Handler for the weaviate acl entries delete operation
func NewWeaviateACLEntriesDelete(ctx *middleware.Context, handler WeaviateACLEntriesDeleteHandler) *WeaviateACLEntriesDelete {
	return &WeaviateACLEntriesDelete{Context: ctx, Handler: handler}
}

/*WeaviateACLEntriesDelete swagger:route DELETE /devices/{deviceId}/aclEntries/{aclEntryId} aclEntries weaviateAclEntriesDelete

Deletes an ACL entry.

*/
type WeaviateACLEntriesDelete struct {
	Context *middleware.Context
	Handler WeaviateACLEntriesDeleteHandler
}

func (o *WeaviateACLEntriesDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateACLEntriesDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
