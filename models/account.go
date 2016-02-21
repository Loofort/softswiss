package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*account Account account

swagger:model account
*/
type Account struct {

	/* money amount

	Required: true
	Minimum: 1
	*/
	Amount float32 `json:"amount,omitempty"`

	/* account id

	Required: true
	*/
	ID int64 `json:"id,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", float32(m.Amount)); err != nil {
		return err
	}

	if err := validate.Minimum("amount", "body", float64(m.Amount), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}
