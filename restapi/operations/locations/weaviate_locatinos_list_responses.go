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
 package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateLocatinosListOKCode is the HTTP code returned for type WeaviateLocatinosListOK
const WeaviateLocatinosListOKCode int = 200

/*WeaviateLocatinosListOK Successful response.

swagger:response weaviateLocatinosListOK
*/
type WeaviateLocatinosListOK struct {

	/*
	  In: Body
	*/
	Payload *models.LocationsListResponse `json:"body,omitempty"`
}

// NewWeaviateLocatinosListOK creates WeaviateLocatinosListOK with default headers values
func NewWeaviateLocatinosListOK() *WeaviateLocatinosListOK {
	return &WeaviateLocatinosListOK{}
}

// WithPayload adds the payload to the weaviate locatinos list o k response
func (o *WeaviateLocatinosListOK) WithPayload(payload *models.LocationsListResponse) *WeaviateLocatinosListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate locatinos list o k response
func (o *WeaviateLocatinosListOK) SetPayload(payload *models.LocationsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateLocatinosListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateLocatinosListNotImplementedCode is the HTTP code returned for type WeaviateLocatinosListNotImplemented
const WeaviateLocatinosListNotImplementedCode int = 501

/*WeaviateLocatinosListNotImplemented Not (yet) implemented.

swagger:response weaviateLocatinosListNotImplemented
*/
type WeaviateLocatinosListNotImplemented struct {
}

// NewWeaviateLocatinosListNotImplemented creates WeaviateLocatinosListNotImplemented with default headers values
func NewWeaviateLocatinosListNotImplemented() *WeaviateLocatinosListNotImplemented {
	return &WeaviateLocatinosListNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateLocatinosListNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
