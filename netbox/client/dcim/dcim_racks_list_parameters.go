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

// NewDcimRacksListParams creates a new DcimRacksListParams object
// with the default values initialized.
func NewDcimRacksListParams() *DcimRacksListParams {
	var ()
	return &DcimRacksListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksListParamsWithTimeout creates a new DcimRacksListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRacksListParamsWithTimeout(timeout time.Duration) *DcimRacksListParams {
	var ()
	return &DcimRacksListParams{

		timeout: timeout,
	}
}

// NewDcimRacksListParamsWithContext creates a new DcimRacksListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRacksListParamsWithContext(ctx context.Context) *DcimRacksListParams {
	var ()
	return &DcimRacksListParams{

		Context: ctx,
	}
}

// NewDcimRacksListParamsWithHTTPClient creates a new DcimRacksListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRacksListParamsWithHTTPClient(client *http.Client) *DcimRacksListParams {
	var ()
	return &DcimRacksListParams{
		HTTPClient: client,
	}
}

/*DcimRacksListParams contains all the parameters to send to the API endpoint
for the dcim racks list operation typically these are written to a http.Request
*/
type DcimRacksListParams struct {

	/*DescUnits*/
	DescUnits *string
	/*FacilityID*/
	FacilityID *string
	/*Group*/
	Group *string
	/*GroupID*/
	GroupID *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *float64
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *string
	/*Serial*/
	Serial *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string
	/*Type*/
	Type *string
	/*UHeight*/
	UHeight *float64
	/*Width*/
	Width *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim racks list params
func (o *DcimRacksListParams) WithTimeout(timeout time.Duration) *DcimRacksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks list params
func (o *DcimRacksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks list params
func (o *DcimRacksListParams) WithContext(ctx context.Context) *DcimRacksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks list params
func (o *DcimRacksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks list params
func (o *DcimRacksListParams) WithHTTPClient(client *http.Client) *DcimRacksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks list params
func (o *DcimRacksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescUnits adds the descUnits to the dcim racks list params
func (o *DcimRacksListParams) WithDescUnits(descUnits *string) *DcimRacksListParams {
	o.SetDescUnits(descUnits)
	return o
}

// SetDescUnits adds the descUnits to the dcim racks list params
func (o *DcimRacksListParams) SetDescUnits(descUnits *string) {
	o.DescUnits = descUnits
}

// WithFacilityID adds the facilityID to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityID(facilityID *string) *DcimRacksListParams {
	o.SetFacilityID(facilityID)
	return o
}

// SetFacilityID adds the facilityId to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityID(facilityID *string) {
	o.FacilityID = facilityID
}

// WithGroup adds the group to the dcim racks list params
func (o *DcimRacksListParams) WithGroup(group *string) *DcimRacksListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the dcim racks list params
func (o *DcimRacksListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupID adds the groupID to the dcim racks list params
func (o *DcimRacksListParams) WithGroupID(groupID *string) *DcimRacksListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the dcim racks list params
func (o *DcimRacksListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithIDIn adds the iDIn to the dcim racks list params
func (o *DcimRacksListParams) WithIDIn(iDIn *float64) *DcimRacksListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim racks list params
func (o *DcimRacksListParams) SetIDIn(iDIn *float64) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the dcim racks list params
func (o *DcimRacksListParams) WithLimit(limit *int64) *DcimRacksListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim racks list params
func (o *DcimRacksListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim racks list params
func (o *DcimRacksListParams) WithName(name *string) *DcimRacksListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim racks list params
func (o *DcimRacksListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim racks list params
func (o *DcimRacksListParams) WithOffset(offset *int64) *DcimRacksListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim racks list params
func (o *DcimRacksListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim racks list params
func (o *DcimRacksListParams) WithQ(q *string) *DcimRacksListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim racks list params
func (o *DcimRacksListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the dcim racks list params
func (o *DcimRacksListParams) WithRole(role *string) *DcimRacksListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the dcim racks list params
func (o *DcimRacksListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the dcim racks list params
func (o *DcimRacksListParams) WithRoleID(roleID *string) *DcimRacksListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the dcim racks list params
func (o *DcimRacksListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithSerial adds the serial to the dcim racks list params
func (o *DcimRacksListParams) WithSerial(serial *string) *DcimRacksListParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the dcim racks list params
func (o *DcimRacksListParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithSite adds the site to the dcim racks list params
func (o *DcimRacksListParams) WithSite(site *string) *DcimRacksListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim racks list params
func (o *DcimRacksListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim racks list params
func (o *DcimRacksListParams) WithSiteID(siteID *string) *DcimRacksListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim racks list params
func (o *DcimRacksListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTenant adds the tenant to the dcim racks list params
func (o *DcimRacksListParams) WithTenant(tenant *string) *DcimRacksListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim racks list params
func (o *DcimRacksListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the dcim racks list params
func (o *DcimRacksListParams) WithTenantID(tenantID *string) *DcimRacksListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim racks list params
func (o *DcimRacksListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithType adds the typeVar to the dcim racks list params
func (o *DcimRacksListParams) WithType(typeVar *string) *DcimRacksListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim racks list params
func (o *DcimRacksListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUHeight adds the uHeight to the dcim racks list params
func (o *DcimRacksListParams) WithUHeight(uHeight *float64) *DcimRacksListParams {
	o.SetUHeight(uHeight)
	return o
}

// SetUHeight adds the uHeight to the dcim racks list params
func (o *DcimRacksListParams) SetUHeight(uHeight *float64) {
	o.UHeight = uHeight
}

// WithWidth adds the width to the dcim racks list params
func (o *DcimRacksListParams) WithWidth(width *string) *DcimRacksListParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the dcim racks list params
func (o *DcimRacksListParams) SetWidth(width *string) {
	o.Width = width
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DescUnits != nil {

		// query param desc_units
		var qrDescUnits string
		if o.DescUnits != nil {
			qrDescUnits = *o.DescUnits
		}
		qDescUnits := qrDescUnits
		if qDescUnits != "" {
			if err := r.SetQueryParam("desc_units", qDescUnits); err != nil {
				return err
			}
		}

	}

	if o.FacilityID != nil {

		// query param facility_id
		var qrFacilityID string
		if o.FacilityID != nil {
			qrFacilityID = *o.FacilityID
		}
		qFacilityID := qrFacilityID
		if qFacilityID != "" {
			if err := r.SetQueryParam("facility_id", qFacilityID); err != nil {
				return err
			}
		}

	}

	if o.Group != nil {

		// query param group
		var qrGroup string
		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {
			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}

	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string
		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {
			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn float64
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := swag.FormatFloat64(qrIDIn)
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
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

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
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

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
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

	if o.UHeight != nil {

		// query param u_height
		var qrUHeight float64
		if o.UHeight != nil {
			qrUHeight = *o.UHeight
		}
		qUHeight := swag.FormatFloat64(qrUHeight)
		if qUHeight != "" {
			if err := r.SetQueryParam("u_height", qUHeight); err != nil {
				return err
			}
		}

	}

	if o.Width != nil {

		// query param width
		var qrWidth string
		if o.Width != nil {
			qrWidth = *o.Width
		}
		qWidth := qrWidth
		if qWidth != "" {
			if err := r.SetQueryParam("width", qWidth); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
