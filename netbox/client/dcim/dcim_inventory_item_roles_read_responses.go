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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimInventoryItemRolesReadReader is a Reader for the DcimInventoryItemRolesRead structure.
type DcimInventoryItemRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemRolesReadOK creates a DcimInventoryItemRolesReadOK with default headers values
func NewDcimInventoryItemRolesReadOK() *DcimInventoryItemRolesReadOK {
	return &DcimInventoryItemRolesReadOK{}
}

/* DcimInventoryItemRolesReadOK describes a response with status code 200, with default header values.

DcimInventoryItemRolesReadOK dcim inventory item roles read o k
*/
type DcimInventoryItemRolesReadOK struct {
	Payload *models.InventoryItemRole
}

func (o *DcimInventoryItemRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesReadOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemRolesReadOK) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesReadDefault creates a DcimInventoryItemRolesReadDefault with default headers values
func NewDcimInventoryItemRolesReadDefault(code int) *DcimInventoryItemRolesReadDefault {
	return &DcimInventoryItemRolesReadDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemRolesReadDefault describes a response with status code -1, with default header values.

DcimInventoryItemRolesReadDefault dcim inventory item roles read default
*/
type DcimInventoryItemRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item roles read default response
func (o *DcimInventoryItemRolesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-item-roles/{id}/][%d] dcim_inventory-item-roles_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
