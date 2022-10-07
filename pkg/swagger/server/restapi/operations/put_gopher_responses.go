// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/scraly/gophers-api/pkg/swagger/server/models"
)

// PutGopherOKCode is the HTTP code returned for type PutGopherOK
const PutGopherOKCode int = 200

/*
PutGopherOK Updated

swagger:response putGopherOK
*/
type PutGopherOK struct {

	/*
	  In: Body
	*/
	Payload *models.Gopher `json:"body,omitempty"`
}

// NewPutGopherOK creates PutGopherOK with default headers values
func NewPutGopherOK() *PutGopherOK {

	return &PutGopherOK{}
}

// WithPayload adds the payload to the put gopher o k response
func (o *PutGopherOK) WithPayload(payload *models.Gopher) *PutGopherOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put gopher o k response
func (o *PutGopherOK) SetPayload(payload *models.Gopher) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutGopherOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutGopherNotFoundCode is the HTTP code returned for type PutGopherNotFound
const PutGopherNotFoundCode int = 404

/*
PutGopherNotFound A gopher with the specified Name was not found.

swagger:response putGopherNotFound
*/
type PutGopherNotFound struct {
}

// NewPutGopherNotFound creates PutGopherNotFound with default headers values
func NewPutGopherNotFound() *PutGopherNotFound {

	return &PutGopherNotFound{}
}

// WriteResponse to the client
func (o *PutGopherNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}