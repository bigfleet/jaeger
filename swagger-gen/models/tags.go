// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Tags Tags
//
// Adds context to a span, for search, viewing and analysis.
//
// For example, a key "your_app.version" would let you lookup traces by version.
// A tag "sql.query" isn't searchable, but it can help in debugging when viewing
// a trace.
//
//
// swagger:model Tags
type Tags map[string]string

// Validate validates this tags
func (m Tags) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tags based on context it is used
func (m Tags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}