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
)

// DcimModulesBulkDeleteReader is a Reader for the DcimModulesBulkDelete structure.
type DcimModulesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModulesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModulesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModulesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModulesBulkDeleteNoContent creates a DcimModulesBulkDeleteNoContent with default headers values
func NewDcimModulesBulkDeleteNoContent() *DcimModulesBulkDeleteNoContent {
	return &DcimModulesBulkDeleteNoContent{}
}

/* DcimModulesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimModulesBulkDeleteNoContent dcim modules bulk delete no content
*/
type DcimModulesBulkDeleteNoContent struct {
}

func (o *DcimModulesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcimModulesBulkDeleteNoContent ", 204)
}

func (o *DcimModulesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModulesBulkDeleteDefault creates a DcimModulesBulkDeleteDefault with default headers values
func NewDcimModulesBulkDeleteDefault(code int) *DcimModulesBulkDeleteDefault {
	return &DcimModulesBulkDeleteDefault{
		_statusCode: code,
	}
}

/* DcimModulesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimModulesBulkDeleteDefault dcim modules bulk delete default
*/
type DcimModulesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim modules bulk delete default response
func (o *DcimModulesBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimModulesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcim_modules_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimModulesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModulesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
