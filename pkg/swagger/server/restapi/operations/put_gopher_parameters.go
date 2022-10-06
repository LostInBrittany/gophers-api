// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
)

// NewPutGopherParams creates a new PutGopherParams object
//
// There are no default values defined in the spec.
func NewPutGopherParams() PutGopherParams {

	return PutGopherParams{}
}

// PutGopherParams contains all the bound params for the put gopher operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutGopher
type PutGopherParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The Gopher to create.
	  In: body
	*/
	Gopher PutGopherBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutGopherParams() beforehand.
func (o *PutGopherParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PutGopherBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("gopher", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Gopher = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
