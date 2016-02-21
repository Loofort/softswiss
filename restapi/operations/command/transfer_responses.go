package command

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/loofort/softswiss/models"
)

/*TransferOK return accounts info with total money amount

swagger:response transferOK
*/
type TransferOK struct {

	// In: body
	Payload []*models.Account `json:"body,omitempty"`
}

// NewTransferOK creates TransferOK with default headers values
func NewTransferOK() *TransferOK {
	return &TransferOK{}
}

// WithPayload adds the payload to the transfer o k response
func (o *TransferOK) WithPayload(payload []*models.Account) *TransferOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *TransferOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*TransferNotFound account not found info

swagger:response transferNotFound
*/
type TransferNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewTransferNotFound creates TransferNotFound with default headers values
func NewTransferNotFound() *TransferNotFound {
	return &TransferNotFound{}
}

// WithPayload adds the payload to the transfer not found response
func (o *TransferNotFound) WithPayload(payload *models.Error) *TransferNotFound {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *TransferNotFound) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TransferDefault unexpected error

swagger:response transferDefault
*/
type TransferDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewTransferDefault creates TransferDefault with default headers values
func NewTransferDefault(code int) *TransferDefault {
	if code <= 0 {
		code = 500
	}

	return &TransferDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the transfer default response
func (o *TransferDefault) WithStatusCode(code int) *TransferDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the transfer default response
func (o *TransferDefault) WithPayload(payload *models.Error) *TransferDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *TransferDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
