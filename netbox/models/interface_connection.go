// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InterfaceConnection interface connection
// swagger:model InterfaceConnection
type InterfaceConnection struct {

	// connection status
	// Required: true
	ConnectionStatus *InterfaceConnectionConnectionStatus `json:"connection_status"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// interface a
	// Required: true
	InterfaceA *PeerInterface `json:"interface_a"`

	// interface b
	// Required: true
	InterfaceB *PeerInterface `json:"interface_b"`
}

// Validate validates this interface connection
func (m *InterfaceConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInterfaceA(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInterfaceB(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceConnection) validateConnectionStatus(formats strfmt.Registry) error {

	if err := validate.Required("connection_status", "body", m.ConnectionStatus); err != nil {
		return err
	}

	if m.ConnectionStatus != nil {

		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceConnection) validateInterfaceA(formats strfmt.Registry) error {

	if err := validate.Required("interface_a", "body", m.InterfaceA); err != nil {
		return err
	}

	if m.InterfaceA != nil {

		if err := m.InterfaceA.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_a")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceConnection) validateInterfaceB(formats strfmt.Registry) error {

	if err := validate.Required("interface_b", "body", m.InterfaceB); err != nil {
		return err
	}

	if m.InterfaceB != nil {

		if err := m.InterfaceB.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_b")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceConnection) UnmarshalBinary(b []byte) error {
	var res InterfaceConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
