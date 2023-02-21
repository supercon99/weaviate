//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Link link
//
// swagger:model Link
type Link struct {
	// weaviate documentation about this resource group
	DocumentationHref string `json:"documentationHref,omitempty"`

	// target of the link
	Href string `json:"href,omitempty"`

	// human readable name of the resource group
	Name string `json:"name,omitempty"`

	// relationship if both resources are related, e.g. 'next', 'previous', 'parent', etc.
	Rel string `json:"rel,omitempty"`
}

// Validate validates this link
func (m *Link) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this link based on context it is used
func (m *Link) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Link) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Link) UnmarshalBinary(b []byte) error {
	var res Link
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
