// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"github.com/go-openapi/analysis/diff"
)

// JSONFormat for json.
const JSONFormat = "json"

// DiffCommand is a command that generates the diff of two swagger specs.
//
// There are no specific options for this expansion.
type DiffCommand struct {
	OnlyBreakingChanges bool   `description:"When present, only shows incompatible changes" long:"break"                                                                        short:"b"`
	Format              string `choice:"txt"                                                choice:"json"                                                                       default:"txt" description:"When present, writes output as json" long:"format" short:"f"`
	IgnoreFile          string `default:"none specified"                                    description:"Exception file of diffs to ignore (copy output from json diff format)" long:"ignore" short:"i"`
	Destination         string `default:"stdout"                                            description:"Output destination file or stdout"                                     long:"dest"   short:"d"`
	Args                struct {
		OldSpec string `positional-arg-name:"{old spec}"`
		NewSpec string `positional-arg-name:"{new spec}"`
	} `required:"2" positional-args:"specs" description:"Input specs to be diff-ed"`
}

// Execute diffs the two specs provided.
func (c *DiffCommand) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

func (c *DiffCommand) readIgnores() (diff.SpecDifferences, error) {
	_ = "STUB: not implemented"
	return *new(diff.SpecDifferences), nil
}

// Open our jsonFile

func (c *DiffCommand) getDiffs() (diff.SpecDifferences, error) {
	_ = "STUB: not implemented"
	return *new(diff.SpecDifferences), nil
}

func (c *DiffCommand) printInfo() { _ = "STUB: not implemented"; return }
