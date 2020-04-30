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

package virtualization

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

	"github.com/cmgreivel/go-netbox/netbox/models"
)

// NewVirtualizationInterfacesCreateParams creates a new VirtualizationInterfacesCreateParams object
// with the default values initialized.
func NewVirtualizationInterfacesCreateParams() *VirtualizationInterfacesCreateParams {
	var ()
	return &VirtualizationInterfacesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesCreateParamsWithTimeout creates a new VirtualizationInterfacesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationInterfacesCreateParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesCreateParams {
	var ()
	return &VirtualizationInterfacesCreateParams{

		timeout: timeout,
	}
}

// NewVirtualizationInterfacesCreateParamsWithContext creates a new VirtualizationInterfacesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationInterfacesCreateParamsWithContext(ctx context.Context) *VirtualizationInterfacesCreateParams {
	var ()
	return &VirtualizationInterfacesCreateParams{

		Context: ctx,
	}
}

// NewVirtualizationInterfacesCreateParamsWithHTTPClient creates a new VirtualizationInterfacesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationInterfacesCreateParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesCreateParams {
	var ()
	return &VirtualizationInterfacesCreateParams{
		HTTPClient: client,
	}
}

/*VirtualizationInterfacesCreateParams contains all the parameters to send to the API endpoint
for the virtualization interfaces create operation typically these are written to a http.Request
*/
type VirtualizationInterfacesCreateParams struct {

	/*Data*/
	Data *models.WritableVirtualMachineInterface

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithContext(ctx context.Context) *VirtualizationInterfacesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithData(data *models.WritableVirtualMachineInterface) *VirtualizationInterfacesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetData(data *models.WritableVirtualMachineInterface) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
