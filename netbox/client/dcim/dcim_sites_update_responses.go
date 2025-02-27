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

// DcimSitesUpdateReader is a Reader for the DcimSitesUpdate structure.
type DcimSitesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesUpdateOK creates a DcimSitesUpdateOK with default headers values
func NewDcimSitesUpdateOK() *DcimSitesUpdateOK {
	return &DcimSitesUpdateOK{}
}

/* DcimSitesUpdateOK describes a response with status code 200, with default header values.

DcimSitesUpdateOK dcim sites update o k
*/
type DcimSitesUpdateOK struct {
	Payload *models.Site
}

func (o *DcimSitesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/sites/{id}/][%d] dcimSitesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSitesUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesUpdateDefault creates a DcimSitesUpdateDefault with default headers values
func NewDcimSitesUpdateDefault(code int) *DcimSitesUpdateDefault {
	return &DcimSitesUpdateDefault{
		_statusCode: code,
	}
}

/* DcimSitesUpdateDefault describes a response with status code -1, with default header values.

DcimSitesUpdateDefault dcim sites update default
*/
type DcimSitesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites update default response
func (o *DcimSitesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSitesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/sites/{id}/][%d] dcim_sites_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSitesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
