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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// CircuitsProviderNetworksCreateReader is a Reader for the CircuitsProviderNetworksCreate structure.
type CircuitsProviderNetworksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProviderNetworksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsProviderNetworksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProviderNetworksCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProviderNetworksCreateCreated creates a CircuitsProviderNetworksCreateCreated with default headers values
func NewCircuitsProviderNetworksCreateCreated() *CircuitsProviderNetworksCreateCreated {
	return &CircuitsProviderNetworksCreateCreated{}
}

/* CircuitsProviderNetworksCreateCreated describes a response with status code 201, with default header values.

CircuitsProviderNetworksCreateCreated circuits provider networks create created
*/
type CircuitsProviderNetworksCreateCreated struct {
	Payload *models.ProviderNetwork
}

func (o *CircuitsProviderNetworksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuitsProviderNetworksCreateCreated  %+v", 201, o.Payload)
}
func (o *CircuitsProviderNetworksCreateCreated) GetPayload() *models.ProviderNetwork {
	return o.Payload
}

func (o *CircuitsProviderNetworksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProviderNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProviderNetworksCreateDefault creates a CircuitsProviderNetworksCreateDefault with default headers values
func NewCircuitsProviderNetworksCreateDefault(code int) *CircuitsProviderNetworksCreateDefault {
	return &CircuitsProviderNetworksCreateDefault{
		_statusCode: code,
	}
}

/* CircuitsProviderNetworksCreateDefault describes a response with status code -1, with default header values.

CircuitsProviderNetworksCreateDefault circuits provider networks create default
*/
type CircuitsProviderNetworksCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits provider networks create default response
func (o *CircuitsProviderNetworksCreateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsProviderNetworksCreateDefault) Error() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuits_provider-networks_create default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsProviderNetworksCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProviderNetworksCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
