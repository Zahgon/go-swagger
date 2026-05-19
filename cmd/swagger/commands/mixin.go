// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"io"

	flags "github.com/jessevdk/go-flags"
)

const (
	// Output messages.
	nothingToDo                           = "nothing to do. Need some swagger files to merge.\nUSAGE: swagger mixin [-c <expected#Collisions>] <primary-swagger-file> <mixin-swagger-file...>"
	ignoreConflictsAndCollisionsSpecified = "both the flags ignore conflicts and collisions were specified. These have conflicting meaning so please only specify one"

	minRequiredMixinArgs = 2
	exitCodeOnCollisions = 254
)

// MixinSpec holds command line flag definitions specific to the mixin
// command. The flags are defined using struct field tags with the
// "github.com/jessevdk/go-flags" format.
type MixinSpec struct {
	ExpectedCollisionCount uint           `description:"expected # of rejected mixin paths, defs, etc due to existing key. Non-zero exit if does not match actual." short:"c"`
	Compact                bool           `description:"applies to JSON formatted specs. When present, doesn't prettify the json"                                   long:"compact"`
	Output                 flags.Filename `description:"the file to write to"                                                                                       long:"output"           short:"o"`
	KeepSpecOrder          bool           `description:"Keep schema properties order identical to spec file"                                                        long:"keep-spec-order"`
	Format                 string         `choice:"yaml"                                                                                                            choice:"json"           default:"json" description:"the format for the spec document" long:"format"`
	IgnoreConflicts        bool           `description:"Ignore conflict"                                                                                            long:"ignore-conflicts"`
}

// Execute runs the mixin command which merges Swagger 2.0 specs into
// one spec
//
// Use cases include adding independently versioned metadata APIs to
// application APIs for microservices.
//
// Typically, multiple APIs to the same service instance is not a
// problem for client generation as you can create more than one
// client to the service from the same calling process (one for each
// API).  However, merging clients can improve clarity of client code
// by having a single client to given service vs several.
//
// Server skeleton generation, ie generating the model & marshaling
// code, http server instance etc. from Swagger, becomes easier with a
// merged spec for some tools & target-languages.  Server code
// generation tools that natively support hosting multiple specs in
// one server process will not need this tool.
func (c *MixinSpec) Execute(args []string) error { _ = "STUB: not implemented"; return nil }

// return the number of unexpected collisions as command exit code.
// Use shell $? to get the actual number of collisions (it has to be non-zero)
// CLI flag value, overflow is not a concern

// return non-zero exit code on merge with collisions

// MixinFiles is a convenience function for Mixin that reads the given
// swagger files, adds the mixins to primary, calls
// FixEmptyResponseDescriptions on the primary, and writes the primary
// with mixins to the given writer in JSON.
//
// Returns the warning messages for collisions that occurred during the mixin
// process and any error.
func (c *MixinSpec) MixinFiles(primaryFile string, mixinFiles []string, _ io.Writer) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
