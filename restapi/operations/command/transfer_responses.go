package command

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/loofort/softswiss/models"
)

/*TransferOK Ok

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

/*TransferDefault transfer default

swagger:response transferDefault
*/
type TransferDefault struct {
	_statusCode int
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

// WriteResponse to the client
func (o *TransferDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
}