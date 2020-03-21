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

// CableLengthUnit Length unit
//
// swagger:model cableLengthUnit
type CableLengthUnit struct {

	// label
	// Required: true
	// Enum: [Meters Centimeters Feet Inches]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [m cm ft in]
	Value *string `json:"value"`
}

// Validate validates this cable length unit
func (m *CableLengthUnit) Validate(formats strfmt.Registry) error {
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

var cableLengthUnitTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Meters","Centimeters","Feet","Inches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableLengthUnitTypeLabelPropEnum = append(cableLengthUnitTypeLabelPropEnum, v)
	}
}

const (

	// CableLengthUnitLabelMeters captures enum value "Meters"
	CableLengthUnitLabelMeters string = "Meters"

	// CableLengthUnitLabelCentimeters captures enum value "Centimeters"
	CableLengthUnitLabelCentimeters string = "Centimeters"

	// CableLengthUnitLabelFeet captures enum value "Feet"
	CableLengthUnitLabelFeet string = "Feet"

	// CableLengthUnitLabelInches captures enum value "Inches"
	CableLengthUnitLabelInches string = "Inches"
)

// prop value enum
func (m *CableLengthUnit) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableLengthUnitTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableLengthUnit) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var cableLengthUnitTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["m","cm","ft","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableLengthUnitTypeValuePropEnum = append(cableLengthUnitTypeValuePropEnum, v)
	}
}

const (

	// CableLengthUnitValueM captures enum value "m"
	CableLengthUnitValueM string = "m"

	// CableLengthUnitValueCm captures enum value "cm"
	CableLengthUnitValueCm string = "cm"

	// CableLengthUnitValueFt captures enum value "ft"
	CableLengthUnitValueFt string = "ft"

	// CableLengthUnitValueIn captures enum value "in"
	CableLengthUnitValueIn string = "in"
)

// prop value enum
func (m *CableLengthUnit) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableLengthUnitTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableLengthUnit) validateValue(formats strfmt.Registry) error {

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
func (m *CableLengthUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableLengthUnit) UnmarshalBinary(b []byte) error {
	var res CableLengthUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
