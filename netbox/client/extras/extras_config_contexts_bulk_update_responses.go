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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// ExtrasConfigContextsBulkUpdateReader is a Reader for the ExtrasConfigContextsBulkUpdate structure.
type ExtrasConfigContextsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasConfigContextsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigContextsBulkUpdateOK creates a ExtrasConfigContextsBulkUpdateOK with default headers values
func NewExtrasConfigContextsBulkUpdateOK() *ExtrasConfigContextsBulkUpdateOK {
	return &ExtrasConfigContextsBulkUpdateOK{}
}

/* ExtrasConfigContextsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsBulkUpdateOK extras config contexts bulk update o k
*/
type ExtrasConfigContextsBulkUpdateOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/config-contexts/][%d] extrasConfigContextsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextsBulkUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigContextsBulkUpdateDefault creates a ExtrasConfigContextsBulkUpdateDefault with default headers values
func NewExtrasConfigContextsBulkUpdateDefault(code int) *ExtrasConfigContextsBulkUpdateDefault {
	return &ExtrasConfigContextsBulkUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasConfigContextsBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasConfigContextsBulkUpdateDefault extras config contexts bulk update default
*/
type ExtrasConfigContextsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras config contexts bulk update default response
func (o *ExtrasConfigContextsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasConfigContextsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/config-contexts/][%d] extras_config-contexts_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasConfigContextsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
