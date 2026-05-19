// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import (
	"github.com/go-swagger/go-swagger/generator"
)

type modelOptions struct {
	ModelPackage               string   `default:"models"                                                                                                    description:"the package to save the models" long:"model-package"   short:"m"`
	Models                     []string `description:"specify a model to include in generation, repeat for multiple (defaults to all)"                       long:"model"                                 short:"M"`
	ExistingModels             string   `description:"use pre-generated models e.g. github.com/foobar/model"                                                 long:"existing-models"`
	StrictAdditionalProperties bool     `description:"disallow extra properties when additionalProperties is set to false"                                   long:"strict-additional-properties"`
	KeepSpecOrder              bool     `description:"keep schema properties order identical to spec file"                                                   long:"keep-spec-order"`
	AllDefinitions             bool     `description:"generate all model definitions regardless of usage in operations"                                      hidden:"deprecated"                          long:"all-definitions"`
	StructTags                 []string `description:"the struct tags to generate, repeat for multiple (defaults to json)"                                   long:"struct-tags"`
	RootedErrorPath            bool     `description:"extends validation errors with the type name instead of an empty path, in the case of arrays and maps" long:"rooted-error-path"`
}

func (mo modelOptions) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

// WithModels adds the model options group.
//
// This group is available to all commands that need some model generation.
type WithModels struct {
	Models modelOptions `group:"Options for model generation"`
}

// Model the generate model file command.
//
// Define the options that are specific to the "swagger generate model" command.
type Model struct {
	WithShared
	WithModels

	NoStruct              bool     `description:"when present will not generate the model struct"                                hidden:"deprecated"            long:"skip-struct"`
	Name                  []string `description:"the model to generate, repeat for multiple (defaults to all). Same as --models" long:"name"                    short:"n"`
	AcceptDefinitionsOnly bool     `description:"accepts a partial swagger spec with only the definitions key"                   long:"accept-definitions-only"`
}

// Execute generates a model file.
func (m *Model) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

// apply options.
func (m Model) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

func (m Model) log(_ string) { _ = "STUB: not implemented"; return }

func (m *Model) generate(opts *generator.GenOpts) error { _ = "STUB: not implemented"; return nil }
