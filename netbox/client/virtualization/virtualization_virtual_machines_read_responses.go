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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmgreivel/go-netbox/netbox/models"
)

// VirtualizationVirtualMachinesReadReader is a Reader for the VirtualizationVirtualMachinesRead structure.
type VirtualizationVirtualMachinesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesReadOK creates a VirtualizationVirtualMachinesReadOK with default headers values
func NewVirtualizationVirtualMachinesReadOK() *VirtualizationVirtualMachinesReadOK {
	return &VirtualizationVirtualMachinesReadOK{}
}

/*VirtualizationVirtualMachinesReadOK handles this case with default header values.

VirtualizationVirtualMachinesReadOK virtualization virtual machines read o k
*/
type VirtualizationVirtualMachinesReadOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesReadOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
