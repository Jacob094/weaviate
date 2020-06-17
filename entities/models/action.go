//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Action action
//
// swagger:model Action
type Action struct {

	// If this object was subject of a classificiation, additional meta info about this classification is available here. (Underscore properties are optional, include them using the ?include=_<propName> parameter)
	Classification *UnderscorePropertiesClassification `json:"_classification,omitempty"`

	// Additional information about how this property was interpreted at vectorization. (Underscore properties are optional, include them using the ?include=_<propName> parameter)
	Interpretation *Interpretation `json:"_interpretation,omitempty"`

	// This object's position in the Contextionary vector space. (Underscore properties are optional, include them using the ?include=_<propName> parameter)
	Vector C11yVector `json:"_vector,omitempty"`

	// Type of the Action, defined in the schema.
	Class string `json:"class,omitempty"`

	// Timestamp of creation of this Action in milliseconds since epoch UTC.
	CreationTimeUnix int64 `json:"creationTimeUnix,omitempty"`

	// ID of the Action.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Timestamp of the last update made to the Action since epoch UTC.
	LastUpdateTimeUnix int64 `json:"lastUpdateTimeUnix,omitempty"`

	// meta
	Meta *UnderscoreProperties `json:"meta,omitempty"`

	// schema
	Schema PropertySchema `json:"schema,omitempty"`

	// vector weights
	VectorWeights VectorWeights `json:"vectorWeights,omitempty"`
}

// Validate validates this action
func (m *Action) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterpretation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Action) validateClassification(formats strfmt.Registry) error {

	if swag.IsZero(m.Classification) { // not required
		return nil
	}

	if m.Classification != nil {
		if err := m.Classification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_classification")
			}
			return err
		}
	}

	return nil
}

func (m *Action) validateInterpretation(formats strfmt.Registry) error {

	if swag.IsZero(m.Interpretation) { // not required
		return nil
	}

	if m.Interpretation != nil {
		if err := m.Interpretation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_interpretation")
			}
			return err
		}
	}

	return nil
}

func (m *Action) validateVector(formats strfmt.Registry) error {

	if swag.IsZero(m.Vector) { // not required
		return nil
	}

	if err := m.Vector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("_vector")
		}
		return err
	}

	return nil
}

func (m *Action) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Action) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Action) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Action) UnmarshalBinary(b []byte) error {
	var res Action
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
