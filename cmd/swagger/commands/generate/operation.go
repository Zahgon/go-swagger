// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import (
	"github.com/go-swagger/go-swagger/generator"
)

type operationOptions struct {
	Operations []string `description:"specify an operation to include, repeat for multiple (defaults to all)" long:"operation"                                 short:"O"`
	Tags       []string `description:"the tags to include, if not specified defaults to all"                  group:"operations"                               long:"tags"`
	APIPackage string   `default:"operations"                                                                 description:"the package to save the operations" long:"api-package" short:"a"`
	WithEnumCI bool     `description:"allow case-insensitive enumerations"                                    long:"with-enum-ci"`

	// tags handling
	SkipTagPackages bool `description:"skips the generation of tag-based operation packages, resulting in a flat generation" long:"skip-tag-packages"`
}

func (oo operationOptions) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

// WithOperations adds the operations options group.
type WithOperations struct {
	Operations operationOptions `group:"Options for operation generation"`
}

// Operation the generate operation files command.
type Operation struct {
	WithShared
	WithOperations

	clientOptions
	serverOptions
	schemeOptions
	mediaOptions

	ModelPackage string `default:"models" description:"the package to save the models" long:"model-package" short:"m"`

	NoHandler    bool `description:"when present will not generate an operation handler"       long:"skip-handler"`
	NoStruct     bool `description:"when present will not generate the parameter model struct" long:"skip-parameters"`
	NoResponses  bool `description:"when present will not generate the response model struct"  long:"skip-responses"`
	NoURLBuilder bool `description:"when present will not generate a URL builder"              long:"skip-url-builder"`

	Name []string `description:"the operations to generate, repeat for multiple (defaults to all). Same as --operations" long:"name" short:"n"`
}

// Execute generates a model file.
func (o *Operation) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

// apply options.
func (o Operation) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

func (o *Operation) generate(opts *generator.GenOpts) error { _ = "STUB: not implemented"; return nil }

func (o Operation) log(_ string) { _ = "STUB: not implemented"; return }
