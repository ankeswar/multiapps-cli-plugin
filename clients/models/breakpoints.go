package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Breakpoints breakpoints
// swagger:model breakpoints
type Breakpoints struct {
	XMLName     xml.Name      `xml:"http://www.sap.com/lmsl/slp breakpoints"`
	Breakpoints []*Breakpoint `xml:"Breakpoint,omitempty"`
}

// Validate validates this breakpoints
func (m Breakpoints) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m.Breakpoints); i++ {

		if swag.IsZero(m.Breakpoints[i]) { // not required
			continue
		}

		if m.Breakpoints[i] != nil {

			if err := m.Breakpoints[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}