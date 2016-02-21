package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/loofort/softswiss/models"
)

/*AccountListOK list the accounts

swagger:response accountListOK
*/
type AccountListOK struct {

	// In: body
	Payload []*models.Account `json:"body,omitempty"`
}

// NewAccountListOK creates AccountListOK with default headers values
func NewAccountListOK() *AccountListOK {
	return &AccountListOK{}
}

// WithPayload adds the payload to the account list o k response
func (o *AccountListOK) WithPayload(payload []*models.Account) *AccountListOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *AccountListOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*AccountListDefault account list default

swagger:response accountListDefault
*/
type AccountListDefault struct {
	_statusCode int
}

// NewAccountListDefault creates AccountListDefault with default headers values
func NewAccountListDefault(code int) *AccountListDefault {
	if code <= 0 {
		code = 500
	}

	return &AccountListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the account list default response
func (o *AccountListDefault) WithStatusCode(code int) *AccountListDefault {
	o._statusCode = code
	return o
}

// WriteResponse to the client
func (o *AccountListDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
}