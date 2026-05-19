package ext

import (
	"context"

	"github.com/go-openapi/strfmt"
)

type MyExtType struct{}

func (MyExtType) Validate(strfmt.Registry) error { _ = "STUB: not implemented"; return nil }
func (MyExtType) ContextValidate(context.Context, strfmt.Registry) error {
	_ = "STUB: not implemented"
	return nil
}
