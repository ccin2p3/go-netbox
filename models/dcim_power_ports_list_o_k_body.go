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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DcimPowerPortsListOKBody dcim power ports list o k body
//
// swagger:model dcimPowerPortsListOKBody
type DcimPowerPortsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*PowerPort `json:"results"`
}

// Validate validates this dcim power ports list o k body
func (m *DcimPowerPortsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DcimPowerPortsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *DcimPowerPortsListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("next", "body", "uri", m.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DcimPowerPortsListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("previous", "body", "uri", m.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DcimPowerPortsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("results", "body", m.Results); err != nil {
		return err
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DcimPowerPortsListOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DcimPowerPortsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimPowerPortsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
