package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

type MtaModules struct {
	Modules []*MtaModulesItems0 `xml:"module,omitempty"`
}

type MtaServices struct {
	Services []string `xml:"service,omitempty"`
}

// Mta mta
// swagger:model mta
type Mta struct {
	XMLName xml.Name `xml:"mta"`

	// metadata
	// Required: true
	Metadata *MtaMetadata `xml:"metadata,omitempty"`

	// modules
	// Required: true
	Modules MtaModules `xml:"modules,omitempty"`

	// services
	// Required: true
	Services MtaServices `xml:"services,omitempty"`
}

// Validate validates this mta
func (m *Mta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mta) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Mta) validateModules(formats strfmt.Registry) error {

	if err := validate.Required("modules", "body", m.Modules); err != nil {
		return err
	}

	for i := 0; i < len(m.Modules.Modules); i++ {

		if swag.IsZero(m.Modules.Modules[i]) { // not required
			continue
		}

		if m.Modules.Modules[i] != nil {

			if err := m.Modules.Modules[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Mta) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	return nil
}

// MtaMetadata mta metadata
// swagger:model MtaMetadata
type MtaMetadata struct {

	// id
	// Required: true
	ID *string `xml:"id"`

	// version
	// Required: true
	Version *string `xml:"version"`
}

// Validate validates this mta metadata
func (m *MtaMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MtaMetadata) validateID(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *MtaMetadata) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

type MtaModulesItems0ProvidedDependencies struct {
	ProvidedDependencies []string `xml:"providedDependency,omitempty"`
}

type MtaModulesItems0Services struct {
	Services []string `xml:"service,omitempty"`
}

// MtaModulesItems0 mta modules items0
// swagger:model MtaModulesItems0
type MtaModulesItems0 struct {

	// app name
	// Required: true
	AppName *string `xml:"appName"`

	// module name
	// Required: true
	ModuleName *string `xml:"moduleName"`

	// provided dependencies
	// Required: true
	ProvidedDependencies MtaModulesItems0ProvidedDependencies `xml:"providedDependencies,omitempty"`

	// services
	// Required: true
	Services MtaModulesItems0Services `xml:"services,omitempty"`
}

// Validate validates this mta modules items0
func (m *MtaModulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModuleName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProvidedDependencies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MtaModulesItems0) validateAppName(formats strfmt.Registry) error {

	if err := validate.Required("appName", "body", m.AppName); err != nil {
		return err
	}

	return nil
}

func (m *MtaModulesItems0) validateModuleName(formats strfmt.Registry) error {

	if err := validate.Required("moduleName", "body", m.ModuleName); err != nil {
		return err
	}

	return nil
}

func (m *MtaModulesItems0) validateProvidedDependencies(formats strfmt.Registry) error {

	if err := validate.Required("providedDependencies", "body", m.ProvidedDependencies); err != nil {
		return err
	}

	return nil
}

func (m *MtaModulesItems0) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	return nil
}