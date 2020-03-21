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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimFrontPortsListParams creates a new DcimFrontPortsListParams object
// with the default values initialized.
func NewDcimFrontPortsListParams() *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsListParamsWithTimeout creates a new DcimFrontPortsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortsListParamsWithTimeout(timeout time.Duration) *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortsListParamsWithContext creates a new DcimFrontPortsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortsListParamsWithContext(ctx context.Context) *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{

		Context: ctx,
	}
}

// NewDcimFrontPortsListParamsWithHTTPClient creates a new DcimFrontPortsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortsListParamsWithHTTPClient(client *http.Client) *DcimFrontPortsListParams {
	var ()
	return &DcimFrontPortsListParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortsListParams contains all the parameters to send to the API endpoint
for the dcim front ports list operation typically these are written to a http.Request
*/
type DcimFrontPortsListParams struct {

	/*Cabled*/
	Cabled *string
	/*Description*/
	Description *string
	/*DescriptionIc*/
	DescriptionIc *string
	/*DescriptionIe*/
	DescriptionIe *string
	/*DescriptionIew*/
	DescriptionIew *string
	/*DescriptionIsw*/
	DescriptionIsw *string
	/*Descriptionn*/
	Descriptionn *string
	/*DescriptionNic*/
	DescriptionNic *string
	/*DescriptionNie*/
	DescriptionNie *string
	/*DescriptionNiew*/
	DescriptionNiew *string
	/*DescriptionNisw*/
	DescriptionNisw *string
	/*Device*/
	Device *string
	/*Devicen*/
	Devicen *string
	/*DeviceID*/
	DeviceID *string
	/*DeviceIDn*/
	DeviceIDn *string
	/*ID*/
	ID *string
	/*IDGt*/
	IDGt *string
	/*IDGte*/
	IDGte *string
	/*IDLt*/
	IDLt *string
	/*IDLte*/
	IDLte *string
	/*IDn*/
	IDn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*NameIc*/
	NameIc *string
	/*NameIe*/
	NameIe *string
	/*NameIew*/
	NameIew *string
	/*NameIsw*/
	NameIsw *string
	/*Namen*/
	Namen *string
	/*NameNic*/
	NameNic *string
	/*NameNie*/
	NameNie *string
	/*NameNiew*/
	NameNiew *string
	/*NameNisw*/
	NameNisw *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Region*/
	Region *string
	/*Regionn*/
	Regionn *string
	/*RegionID*/
	RegionID *string
	/*RegionIDn*/
	RegionIDn *string
	/*Site*/
	Site *string
	/*Siten*/
	Siten *string
	/*SiteID*/
	SiteID *string
	/*SiteIDn*/
	SiteIDn *string
	/*Tag*/
	Tag *string
	/*Tagn*/
	Tagn *string
	/*Type*/
	Type *string
	/*Typen*/
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithTimeout(timeout time.Duration) *DcimFrontPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithContext(ctx context.Context) *DcimFrontPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithHTTPClient(client *http.Client) *DcimFrontPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithCabled(cabled *string) *DcimFrontPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithDescription adds the description to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescription(description *string) *DcimFrontPortsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionIc(descriptionIc *string) *DcimFrontPortsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionIe(descriptionIe *string) *DcimFrontPortsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionIew(descriptionIew *string) *DcimFrontPortsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionIsw(descriptionIsw *string) *DcimFrontPortsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionn(descriptionn *string) *DcimFrontPortsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionNic(descriptionNic *string) *DcimFrontPortsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionNie(descriptionNie *string) *DcimFrontPortsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionNiew(descriptionNiew *string) *DcimFrontPortsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDescriptionNisw(descriptionNisw *string) *DcimFrontPortsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithDevice adds the device to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDevice(device *string) *DcimFrontPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDevicen(devicen *string) *DcimFrontPortsListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDeviceID(deviceID *string) *DcimFrontPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithDeviceIDn(deviceIDn *string) *DcimFrontPortsListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithID adds the id to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithID(id *string) *DcimFrontPortsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithIDGt(iDGt *string) *DcimFrontPortsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithIDGte(iDGte *string) *DcimFrontPortsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithIDLt(iDLt *string) *DcimFrontPortsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithIDLte(iDLte *string) *DcimFrontPortsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithIDn(iDn *string) *DcimFrontPortsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithLimit(limit *int64) *DcimFrontPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithName(name *string) *DcimFrontPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameIc(nameIc *string) *DcimFrontPortsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameIe(nameIe *string) *DcimFrontPortsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameIew(nameIew *string) *DcimFrontPortsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameIsw(nameIsw *string) *DcimFrontPortsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNamen(namen *string) *DcimFrontPortsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameNic(nameNic *string) *DcimFrontPortsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameNie(nameNie *string) *DcimFrontPortsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameNiew(nameNiew *string) *DcimFrontPortsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithNameNisw(nameNisw *string) *DcimFrontPortsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithOffset(offset *int64) *DcimFrontPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithQ(q *string) *DcimFrontPortsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithRegion(region *string) *DcimFrontPortsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithRegionn(regionn *string) *DcimFrontPortsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithRegionID(regionID *string) *DcimFrontPortsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithRegionIDn(regionIDn *string) *DcimFrontPortsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithSite(site *string) *DcimFrontPortsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithSiten(siten *string) *DcimFrontPortsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithSiteID(siteID *string) *DcimFrontPortsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithSiteIDn(siteIDn *string) *DcimFrontPortsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithTag(tag *string) *DcimFrontPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithTagn(tagn *string) *DcimFrontPortsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithType adds the typeVar to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithType(typeVar *string) *DcimFrontPortsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim front ports list params
func (o *DcimFrontPortsListParams) WithTypen(typen *string) *DcimFrontPortsListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim front ports list params
func (o *DcimFrontPortsListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string
		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {
			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}

	}

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string
		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {
			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string
		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {
			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string
		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {
			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string
		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {
			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}

	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string
		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {
			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string
		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {
			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string
		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {
			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string
		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {
			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}

	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string
		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {
			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
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

	if o.Devicen != nil {

		// query param device__n
		var qrDevicen string
		if o.Devicen != nil {
			qrDevicen = *o.Devicen
		}
		qDevicen := qrDevicen
		if qDevicen != "" {
			if err := r.SetQueryParam("device__n", qDevicen); err != nil {
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

	if o.DeviceIDn != nil {

		// query param device_id__n
		var qrDeviceIDn string
		if o.DeviceIDn != nil {
			qrDeviceIDn = *o.DeviceIDn
		}
		qDeviceIDn := qrDeviceIDn
		if qDeviceIDn != "" {
			if err := r.SetQueryParam("device_id__n", qDeviceIDn); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string
		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {
			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}

	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string
		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {
			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}

	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string
		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {
			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}

	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string
		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {
			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}

	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string
		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {
			if err := r.SetQueryParam("id__n", qIDn); err != nil {
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

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string
		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {
			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}

	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string
		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {
			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}

	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string
		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {
			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}

	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string
		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {
			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}

	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string
		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {
			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}

	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string
		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {
			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}

	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string
		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {
			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}

	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string
		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {
			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}

	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string
		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {
			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
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

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.Regionn != nil {

		// query param region__n
		var qrRegionn string
		if o.Regionn != nil {
			qrRegionn = *o.Regionn
		}
		qRegionn := qrRegionn
		if qRegionn != "" {
			if err := r.SetQueryParam("region__n", qRegionn); err != nil {
				return err
			}
		}

	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {
			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}

	}

	if o.RegionIDn != nil {

		// query param region_id__n
		var qrRegionIDn string
		if o.RegionIDn != nil {
			qrRegionIDn = *o.RegionIDn
		}
		qRegionIDn := qrRegionIDn
		if qRegionIDn != "" {
			if err := r.SetQueryParam("region_id__n", qRegionIDn); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.Siten != nil {

		// query param site__n
		var qrSiten string
		if o.Siten != nil {
			qrSiten = *o.Siten
		}
		qSiten := qrSiten
		if qSiten != "" {
			if err := r.SetQueryParam("site__n", qSiten); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.SiteIDn != nil {

		// query param site_id__n
		var qrSiteIDn string
		if o.SiteIDn != nil {
			qrSiteIDn = *o.SiteIDn
		}
		qSiteIDn := qrSiteIDn
		if qSiteIDn != "" {
			if err := r.SetQueryParam("site_id__n", qSiteIDn); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string
		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {
			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string
		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {
			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
