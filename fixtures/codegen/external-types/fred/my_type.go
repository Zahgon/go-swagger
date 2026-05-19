package fred

import (
	"context"
	"io"

	"github.com/go-openapi/strfmt"
)

// MyAlternateType ...
type MyAlternateType string

// Validate MyAlternateType
func (MyAlternateType) Validate(strfmt.Registry) error { _ = "STUB: not implemented"; return nil }
func (MyAlternateType) ContextValidate(context.Context, strfmt.Registry) error {
	_ = "STUB: not implemented"

	// MyAlternateInteger ...
	return nil
}

type MyAlternateInteger int

// Validate MyAlternateInteger
func (MyAlternateInteger) Validate(strfmt.Registry) error { _ = "STUB: not implemented"; return nil }
func (MyAlternateInteger) ContextValidate(context.Context, strfmt.Registry) error {
	_ = "STUB: not implemented"

	// MyAlternateString ...
	return nil
}

type MyAlternateString string

// Validate MyAlternateString
func (MyAlternateString) Validate(strfmt.Registry) error { _ = "STUB: not implemented"; return nil }
func (MyAlternateString) ContextValidate(context.Context, strfmt.Registry) error {
	_ = "STUB: not implemented"

	// MyAlternateOtherType ...
	return nil
}

type MyAlternateOtherType struct{}

// Validate MyAlternateOtherType
func (MyAlternateOtherType) Validate(strfmt.Registry) error { _ = "STUB: not implemented"; return nil }
func (MyAlternateOtherType) ContextValidate(context.Context, strfmt.Registry) error {
	_ = "STUB: not implemented"

	// MyAlternateStreamer ...
	return nil
}

type MyAlternateStreamer io.Reader

// MyAlternateInterface ...
type MyAlternateInterface any
