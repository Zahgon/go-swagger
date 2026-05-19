// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import "github.com/go-swagger/go-swagger/generator"

type Cli struct {
	// generate a cli includes all client code
	Client

	// cmd/<cli-app-name>/main.go will be generated. This ensures that go install will compile the app with desired name.
	CliAppName string `default:"cli" description:"the app name for the cli executable. useful for go install." long:"cli-app-name"`
	CliPackage string `default:"cli" description:"the package to save the cli specific code"                   long:"cli-package"`
}

// Execute runs this command.
func (c *Cli) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

// apply options.
func (c Cli) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

func (c *Cli) generate(opts *generator.GenOpts) error { _ = "STUB: not implemented"; return nil }
