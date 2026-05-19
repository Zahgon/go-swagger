// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"io"
	"io/fs"
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/go-openapi/spec"
)

const readableMode fs.FileMode = 0o644 & fs.ModePerm

// ExpandSpec is a command that expands the $refs in a swagger document.
//
// There are no specific options for this expansion.
type ExpandSpec struct {
	Compact bool           `description:"applies to JSON formatted specs. When present, doesn't prettify the json" long:"compact"`
	Output  flags.Filename `description:"the file to write to"                                                     long:"output"  short:"o"`
	Format  string         `choice:"yaml"                                                                          choice:"json"  default:"json" description:"the format for the spec document" long:"format"`
}

// Execute expands the spec.
func (c *ExpandSpec) Execute(args []string) error { _ = "STUB: not implemented"; return nil }

var defaultWriter io.Writer = os.Stdout

func writeToFile(swspec *spec.Swagger, pretty bool, format string, output string) error {
	_ = "STUB: not implemented"
	return nil
}

func marshalAsYAML(swspec *spec.Swagger) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
