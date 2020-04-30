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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cmgreivel/go-netbox/netbox/models"
)

// NewCircuitsProvidersUpdateParams creates a new CircuitsProvidersUpdateParams object
// with the default values initialized.
func NewCircuitsProvidersUpdateParams() *CircuitsProvidersUpdateParams {
	var ()
	return &CircuitsProvidersUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersUpdateParamsWithTimeout creates a new CircuitsProvidersUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsProvidersUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersUpdateParams {
	var ()
	return &CircuitsProvidersUpdateParams{

		timeout: timeout,
	}
}

// NewCircuitsProvidersUpdateParamsWithContext creates a new CircuitsProvidersUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsProvidersUpdateParamsWithContext(ctx context.Context) *CircuitsProvidersUpdateParams {
	var ()
	return &CircuitsProvidersUpdateParams{

		Context: ctx,
	}
}

// NewCircuitsProvidersUpdateParamsWithHTTPClient creates a new CircuitsProvidersUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsProvidersUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersUpdateParams {
	var ()
	return &CircuitsProvidersUpdateParams{
		HTTPClient: client,
	}
}

/*CircuitsProvidersUpdateParams contains all the parameters to send to the API endpoint
for the circuits providers update operation typically these are written to a http.Request
*/
type CircuitsProvidersUpdateParams struct {

	/*Data*/
	Data *models.Provider
	/*ID
	  A unique integer value identifying this provider.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithContext(ctx context.Context) *CircuitsProvidersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithData(data *models.Provider) *CircuitsProvidersUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WithID adds the id to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) WithID(id int64) *CircuitsProvidersUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits providers update params
func (o *CircuitsProvidersUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
