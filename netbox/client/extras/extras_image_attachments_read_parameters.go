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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewExtrasImageAttachmentsReadParams creates a new ExtrasImageAttachmentsReadParams object
// with the default values initialized.
func NewExtrasImageAttachmentsReadParams() *ExtrasImageAttachmentsReadParams {
	var ()
	return &ExtrasImageAttachmentsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsReadParamsWithTimeout creates a new ExtrasImageAttachmentsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasImageAttachmentsReadParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsReadParams {
	var ()
	return &ExtrasImageAttachmentsReadParams{

		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsReadParamsWithContext creates a new ExtrasImageAttachmentsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasImageAttachmentsReadParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsReadParams {
	var ()
	return &ExtrasImageAttachmentsReadParams{

		Context: ctx,
	}
}

// NewExtrasImageAttachmentsReadParamsWithHTTPClient creates a new ExtrasImageAttachmentsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasImageAttachmentsReadParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsReadParams {
	var ()
	return &ExtrasImageAttachmentsReadParams{
		HTTPClient: client,
	}
}

/*ExtrasImageAttachmentsReadParams contains all the parameters to send to the API endpoint
for the extras image attachments read operation typically these are written to a http.Request
*/
type ExtrasImageAttachmentsReadParams struct {

	/*ID
	  A unique integer value identifying this image attachment.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) WithID(id int64) *ExtrasImageAttachmentsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments read params
func (o *ExtrasImageAttachmentsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
