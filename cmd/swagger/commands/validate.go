// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package commands

const (
	// Output messages.
	missingArgMsg  = "the validate command requires the swagger document url to be specified"
	validSpecMsg   = "\nThe swagger spec at %q is valid against swagger specification %s\n"
	invalidSpecMsg = "\nThe swagger spec at %q is invalid against swagger specification %s.\nSee errors below:\n"
	warningSpecMsg = "\nThe swagger spec at %q showed up some valid but possibly unwanted constructs."
)

// ValidateSpec is a command that validates a swagger document
// against the swagger specification.
type ValidateSpec struct {
	// SchemaURL string `long:"schema" description:"The schema url to use" default:"http://swagger.io/v2/schema.json"`
	SkipWarnings bool `description:"when present will not show up warnings upon validation"                    long:"skip-warnings"`
	StopOnError  bool `description:"when present will not continue validation after critical errors are found" long:"stop-on-error"`
}

// Execute validates the spec.
func (c *ValidateSpec) Execute(args []string) error { _ = "STUB: not implemented"; return nil }

// Attempts to report about all errors

// returns fully detailed result with errors and warnings
