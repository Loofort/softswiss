package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/loofort/softswiss/models"
)

/*AccountRegistartionCreated Created

swagger:response accountRegistartionCreated
*/
type AccountRegistartionCreated struct {

	// In: body
	Payload *models.Account `json:"body,omitempty"`
}

// NewAccountRegistartionCreated creates AccountRegistartionCreated with default headers values
func NewAccountRegistartionCreated() *AccountRegistartionCreated {
	return &AccountRegistartionCreated{}
}

// WithPayload adds the payload to the account registartion created response
func (o *AccountRegistartionCreated) WithPayload(payload *models.Account) *AccountRegistartionCreated {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AccountRegistartionCreated) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AccountRegistartionDefault unexpected error

swagger:response accountRegistartionDefault
*/
type AccountRegistartionDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewAccountRegistartionDefault creates AccountRegistartionDefault with default headers values
func NewAccountRegistartionDefault(code int) *AccountRegistartionDefault {
	if code <= 0 {
		code = 500
	}

	return &AccountRegistartionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the account registartion default response
func (o *AccountRegistartionDefault) WithStatusCode(code int) *AccountRegistartionDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the account registartion default response
func (o *AccountRegistartionDefault) WithPayload(payload *models.Error) *AccountRegistartionDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AccountRegistartionDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
