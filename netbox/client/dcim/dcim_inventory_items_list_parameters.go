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

package dcim

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

// NewDcimInventoryItemsListParams creates a new DcimInventoryItemsListParams object
// with the default values initialized.
func NewDcimInventoryItemsListParams() *DcimInventoryItemsListParams {
	var ()
	return &DcimInventoryItemsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsListParamsWithTimeout creates a new DcimInventoryItemsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInventoryItemsListParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsListParams {
	var ()
	return &DcimInventoryItemsListParams{

		timeout: timeout,
	}
}

// NewDcimInventoryItemsListParamsWithContext creates a new DcimInventoryItemsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInventoryItemsListParamsWithContext(ctx context.Context) *DcimInventoryItemsListParams {
	var ()
	return &DcimInventoryItemsListParams{

		Context: ctx,
	}
}

// NewDcimInventoryItemsListParamsWithHTTPClient creates a new DcimInventoryItemsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInventoryItemsListParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsListParams {
	var ()
	return &DcimInventoryItemsListParams{
		HTTPClient: client,
	}
}

/*DcimInventoryItemsListParams contains all the parameters to send to the API endpoint
for the dcim inventory items list operation typically these are written to a http.Request
*/
type DcimInventoryItemsListParams struct {

	/*AssetTag*/
	AssetTag *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
	/*Discovered*/
	Discovered *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Manufacturer*/
	Manufacturer *string
	/*ManufacturerID*/
	ManufacturerID *string
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*ParentID*/
	ParentID *string
	/*PartID*/
	PartID *string
	/*Q*/
	Q *string
	/*Serial*/
	Serial *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithContext(ctx context.Context) *DcimInventoryItemsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetTag adds the assetTag to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTag(assetTag *string) *DcimInventoryItemsListParams {
	o.SetAssetTag(assetTag)
	return o
}

// SetAssetTag adds the assetTag to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTag(assetTag *string) {
	o.AssetTag = assetTag
}

// WithDevice adds the device to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDevice(device *string) *DcimInventoryItemsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDeviceID(deviceID *string) *DcimInventoryItemsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDiscovered adds the discovered to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDiscovered(discovered *string) *DcimInventoryItemsListParams {
	o.SetDiscovered(discovered)
	return o
}

// SetDiscovered adds the discovered to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDiscovered(discovered *string) {
	o.Discovered = discovered
}

// WithLimit adds the limit to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithLimit(limit *int64) *DcimInventoryItemsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithManufacturer(manufacturer *string) *DcimInventoryItemsListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturerID adds the manufacturerID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithManufacturerID(manufacturerID *string) *DcimInventoryItemsListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithName adds the name to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithName(name *string) *DcimInventoryItemsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithOffset(offset *int64) *DcimInventoryItemsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParentID adds the parentID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithParentID(parentID *string) *DcimInventoryItemsListParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetParentID(parentID *string) {
	o.ParentID = parentID
}

// WithPartID adds the partID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartID(partID *string) *DcimInventoryItemsListParams {
	o.SetPartID(partID)
	return o
}

// SetPartID adds the partId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartID(partID *string) {
	o.PartID = partID
}

// WithQ adds the q to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithQ(q *string) *DcimInventoryItemsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSerial adds the serial to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithSerial(serial *string) *DcimInventoryItemsListParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssetTag != nil {

		// query param asset_tag
		var qrAssetTag string
		if o.AssetTag != nil {
			qrAssetTag = *o.AssetTag
		}
		qAssetTag := qrAssetTag
		if qAssetTag != "" {
			if err := r.SetQueryParam("asset_tag", qAssetTag); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.Discovered != nil {

		// query param discovered
		var qrDiscovered string
		if o.Discovered != nil {
			qrDiscovered = *o.Discovered
		}
		qDiscovered := qrDiscovered
		if qDiscovered != "" {
			if err := r.SetQueryParam("discovered", qDiscovered); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string
		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {
			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
				return err
			}
		}

	}

	if o.ManufacturerID != nil {

		// query param manufacturer_id
		var qrManufacturerID string
		if o.ManufacturerID != nil {
			qrManufacturerID = *o.ManufacturerID
		}
		qManufacturerID := qrManufacturerID
		if qManufacturerID != "" {
			if err := r.SetQueryParam("manufacturer_id", qManufacturerID); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.ParentID != nil {

		// query param parent_id
		var qrParentID string
		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := qrParentID
		if qParentID != "" {
			if err := r.SetQueryParam("parent_id", qParentID); err != nil {
				return err
			}
		}

	}

	if o.PartID != nil {

		// query param part_id
		var qrPartID string
		if o.PartID != nil {
			qrPartID = *o.PartID
		}
		qPartID := qrPartID
		if qPartID != "" {
			if err := r.SetQueryParam("part_id", qPartID); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Serial != nil {

		// query param serial
		var qrSerial string
		if o.Serial != nil {
			qrSerial = *o.Serial
		}
		qSerial := qrSerial
		if qSerial != "" {
			if err := r.SetQueryParam("serial", qSerial); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
