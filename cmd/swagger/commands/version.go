// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package commands

var (
	// Version for the swagger command.
	Version string
	// Commit for the swagger command.
	Commit string
)

// PrintVersion the command.
type PrintVersion struct{}

// Execute this command.
//
//nolint:forbidigo // this commands is allowed to use fmt.Println
func (p *PrintVersion) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

// built from source, with module (e.g. go get)

// built from source, local repo

// released version
