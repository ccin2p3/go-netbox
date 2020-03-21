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

// PrefixStatus Status
//
// swagger:model prefixStatus
type PrefixStatus struct {

	// label
	// Required: true
	// Enum: [Container Active Reserved Deprecated]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [container active reserved deprecated]
	Value *string `json:"value"`
}

// Validate validates this prefix status
func (m *PrefixStatus) Validate(formats strfmt.Registry) error {
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

var prefixStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Container","Active","Reserved","Deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		prefixStatusTypeLabelPropEnum = append(prefixStatusTypeLabelPropEnum, v)
	}
}

const (

	// PrefixStatusLabelContainer captures enum value "Container"
	PrefixStatusLabelContainer string = "Container"

	// PrefixStatusLabelActive captures enum value "Active"
	PrefixStatusLabelActive string = "Active"

	// PrefixStatusLabelReserved captures enum value "Reserved"
	PrefixStatusLabelReserved string = "Reserved"

	// PrefixStatusLabelDeprecated captures enum value "Deprecated"
	PrefixStatusLabelDeprecated string = "Deprecated"
)

// prop value enum
func (m *PrefixStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, prefixStatusTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PrefixStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var prefixStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["container","active","reserved","deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		prefixStatusTypeValuePropEnum = append(prefixStatusTypeValuePropEnum, v)
	}
}

const (

	// PrefixStatusValueContainer captures enum value "container"
	PrefixStatusValueContainer string = "container"

	// PrefixStatusValueActive captures enum value "active"
	PrefixStatusValueActive string = "active"

	// PrefixStatusValueReserved captures enum value "reserved"
	PrefixStatusValueReserved string = "reserved"

	// PrefixStatusValueDeprecated captures enum value "deprecated"
	PrefixStatusValueDeprecated string = "deprecated"
)

// prop value enum
func (m *PrefixStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, prefixStatusTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PrefixStatus) validateValue(formats strfmt.Registry) error {

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
func (m *PrefixStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrefixStatus) UnmarshalBinary(b []byte) error {
	var res PrefixStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
