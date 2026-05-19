// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import (
	"io"
	"os"

	"github.com/go-openapi/spec"

	"github.com/jessevdk/go-flags"
)

const (
	allFromCurrent = "./..."
)

// SpecFile command to generate a swagger spec from a go application.
type SpecFile struct {
	WorkDir                 string         `default:"."                                                                                                  description:"the base path to use" long:"work-dir" short:"w"`
	BuildTags               string         `default:""                                                                                                   description:"build tags"           long:"tags"     short:"t"`
	ScanModels              bool           `description:"includes models that were annotated with 'swagger:model'"                                       long:"scan-models"                 short:"m"`
	Compact                 bool           `description:"when present, doesn't prettify the json"                                                        long:"compact"`
	Output                  flags.Filename `description:"the file to write to"                                                                           long:"output"                      short:"o"`
	Input                   flags.Filename `description:"an input swagger file with which to merge"                                                      long:"input"                       short:"i"`
	Include                 []string       `description:"include packages matching pattern"                                                              long:"include"                     short:"c"`
	Exclude                 []string       `description:"exclude packages matching pattern"                                                              long:"exclude"                     short:"x"`
	IncludeTags             []string       `description:"include routes having specified tags (can be specified many times)"                             long:"include-tag"                 short:""`
	ExcludeTags             []string       `description:"exclude routes having specified tags (can be specified many times)"                             long:"exclude-tag"                 short:""`
	ExcludeDeps             bool           `description:"exclude all dependencies of project"                                                            long:"exclude-deps"                short:""`
	SetXNullableForPointers bool           `description:"set x-nullable extension to true automatically for fields of pointer types without 'omitempty'" long:"nullable-pointers"           short:"n"`
	RefAliases              bool           `description:"transform aliased types into $ref rather than expanding their definition"                       long:"ref-aliases"                 short:"r"`
	TransparentAliases      bool           `description:"treat type aliases as completely transparent, never creating definitions for them"              long:"transparent-aliases"         short:""`
	DescWithRef             bool           `description:"allow descriptions to flow alongside $ref"                                                      long:"allow-desc-with-ref"         short:""`
	Format                  string         `choice:"yaml"                                                                                                choice:"json"                      default:"json"  description:"the format for the spec document" long:"format"`
}

// Execute runs this command.
func (s *SpecFile) Execute(args []string) error {
	_ = "STUB: not implemented"
	// by default consider all the paths under the working directory
	return nil
}

// load an external spec to merge into

func loadSpec(input string) (*spec.Swagger, error) { _ = "STUB: not implemented"; return nil, nil }

var defaultWriter io.Writer = os.Stdout

const generatedFileMode os.FileMode = 0o644

func writeToFile(swspec *spec.Swagger, pretty bool, format string, output string) error {
	_ = "STUB: not implemented"
	return nil
}

//#nosec

// #nosec

func marshalToJSONFormat(swspec *spec.Swagger, pretty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalToYAMLFormat(swspec *spec.Swagger) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
