// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InventoryListAgentsResponse inventory list agents response
// swagger:model inventoryListAgentsResponse
type InventoryListAgentsResponse struct {

	// mysqld exporter
	MysqldExporter []*InventoryMySqldExporter `json:"mysqld_exporter"`
}

// Validate validates this inventory list agents response
func (m *InventoryListAgentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryListAgentsResponse) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(m.MysqldExporter) { // not required
		return nil
	}

	for i := 0; i < len(m.MysqldExporter); i++ {
		if swag.IsZero(m.MysqldExporter[i]) { // not required
			continue
		}

		if m.MysqldExporter[i] != nil {
			if err := m.MysqldExporter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mysqld_exporter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryListAgentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryListAgentsResponse) UnmarshalBinary(b []byte) error {
	var res InventoryListAgentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}