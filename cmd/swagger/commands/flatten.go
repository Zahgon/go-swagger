// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	flags "github.com/jessevdk/go-flags"

	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
)

// FlattenSpec is a command that flattens a swagger document
// which will expand the remote references in a spec and move inline schemas to definitions
// after flattening there are no complex inlined anymore.
type FlattenSpec struct {
	generate.FlattenCmdOptions

	Compact bool           `description:"applies to JSON formatted specs. When present, doesn't prettify the json" long:"compact"`
	Output  flags.Filename `description:"the file to write to"                                                     long:"output"  short:"o"`
	Format  string         `choice:"yaml"                                                                          choice:"json"  default:"json" description:"the format for the spec document" long:"format"`
}

// Execute flattens the spec.
func (c *FlattenSpec) Execute(args []string) error { _ = "STUB: not implemented"; return nil }

// defaults
