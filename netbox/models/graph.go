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

// Graph graph
// swagger:model Graph
type Graph struct {

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Link URL
	// Max Length: 200
	Link strfmt.URI `json:"link,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	Name *string `json:"name"`

	// Source URL
	// Required: true
	// Max Length: 500
	Source *string `json:"source"`

	// type
	// Required: true
	Type *GraphType `json:"type"`

	// Weight
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this graph
func (m *Graph) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Graph) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if err := validate.MaxLength("link", "body", string(m.Link), 200); err != nil {
		return err
	}

	if err := validate.FormatOf("link", "body", "uri", m.Link.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Graph) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *Graph) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if err := validate.MaxLength("source", "body", string(*m.Source), 500); err != nil {
		return err
	}

	return nil
}

func (m *Graph) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {

		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *Graph) validateWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", int64(*m.Weight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", int64(*m.Weight), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Graph) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Graph) UnmarshalBinary(b []byte) error {
	var res Graph
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
