// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Createjob createjob
// swagger:model createjob
type Createjob struct {

	// compliancetype
	// Required: true
	// Min Length: 1
	Compliancetype *string `json:"compliancetype"`

	// hostname
	// Required: true
	// Min Length: 1
	Hostname *string `json:"hostname"`

	// id
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// password
	// Required: true
	// Min Length: 1
	Password *string `json:"password"`

	// username
	// Required: true
	// Min Length: 1
	Username *string `json:"username"`
}

// Validate validates this createjob
func (m *Createjob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompliancetype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Createjob) validateCompliancetype(formats strfmt.Registry) error {

	if err := validate.Required("compliancetype", "body", m.Compliancetype); err != nil {
		return err
	}

	if err := validate.MinLength("compliancetype", "body", string(*m.Compliancetype), 1); err != nil {
		return err
	}

	return nil
}

func (m *Createjob) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	if err := validate.MinLength("hostname", "body", string(*m.Hostname), 1); err != nil {
		return err
	}

	return nil
}

func (m *Createjob) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *Createjob) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 1); err != nil {
		return err
	}

	return nil
}

func (m *Createjob) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", string(*m.Username), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Createjob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Createjob) UnmarshalBinary(b []byte) error {
	var res Createjob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
