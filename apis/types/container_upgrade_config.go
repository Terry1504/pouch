// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerUpgradeConfig ContainerUpgradeConfig is used for API "POST /containers/{name:.*}/upgrade". when upgrade a container,
// we must specify new image used to create a new container, and also can specify `Cmd` and `Entrypoint` for
// new container. There is all parameters that upgrade a container, if want to change other parameters, i
// think you should use `update` API interface.
//
// swagger:model ContainerUpgradeConfig

type ContainerUpgradeConfig struct {

	// Execution commands and args
	Cmd []string `json:"Cmd"`

	// The entrypoint for the container as a string or an array of strings.
	// If the array consists of exactly one empty string (`[""]`) then the entry point is reset to system default.
	//
	Entrypoint []string `json:"Entrypoint"`

	// image
	// Required: true
	Image string `json:"Image"`
}

/* polymorph ContainerUpgradeConfig Cmd false */

/* polymorph ContainerUpgradeConfig Entrypoint false */

/* polymorph ContainerUpgradeConfig Image false */

// Validate validates this container upgrade config
func (m *ContainerUpgradeConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCmd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEntrypoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerUpgradeConfig) validateCmd(formats strfmt.Registry) error {

	if swag.IsZero(m.Cmd) { // not required
		return nil
	}

	return nil
}

func (m *ContainerUpgradeConfig) validateEntrypoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Entrypoint) { // not required
		return nil
	}

	return nil
}

func (m *ContainerUpgradeConfig) validateImage(formats strfmt.Registry) error {

	if err := validate.RequiredString("Image", "body", string(m.Image)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerUpgradeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerUpgradeConfig) UnmarshalBinary(b []byte) error {
	var res ContainerUpgradeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
