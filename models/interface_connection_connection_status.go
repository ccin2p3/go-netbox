// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
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
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InterfaceConnectionConnectionStatus Connection status
//
// swagger:model interfaceConnectionConnectionStatus
type InterfaceConnectionConnectionStatus struct {

	// label
	// Required: true
	// Enum: [Not Connected Connected]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [false true]
	Value *int64 `json:"value"`
}

// Validate validates this interface connection connection status
func (m *InterfaceConnectionConnectionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var interfaceConnectionConnectionStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Not Connected","Connected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceConnectionConnectionStatusTypeLabelPropEnum = append(interfaceConnectionConnectionStatusTypeLabelPropEnum, v)
	}
}

const (

	// InterfaceConnectionConnectionStatusLabelNotConnected captures enum value "Not Connected"
	InterfaceConnectionConnectionStatusLabelNotConnected string = "Not Connected"

	// InterfaceConnectionConnectionStatusLabelConnected captures enum value "Connected"
	InterfaceConnectionConnectionStatusLabelConnected string = "Connected"
)

// prop value enum
func (m *InterfaceConnectionConnectionStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, interfaceConnectionConnectionStatusTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceConnectionConnectionStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var interfaceConnectionConnectionStatusTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceConnectionConnectionStatusTypeValuePropEnum = append(interfaceConnectionConnectionStatusTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *InterfaceConnectionConnectionStatus) validateValueEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, interfaceConnectionConnectionStatusTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceConnectionConnectionStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceConnectionConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceConnectionConnectionStatus) UnmarshalBinary(b []byte) error {
	var res InterfaceConnectionConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
