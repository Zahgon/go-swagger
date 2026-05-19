// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import (
	"github.com/go-swagger/go-swagger/generator"
)

// Support generates the supporting files.
type Support struct {
	WithShared
	WithModels
	WithOperations

	clientOptions
	serverOptions
	schemeOptions
	mediaOptions

	Name string `description:"the name of the application, defaults to a mangled value of info.title" long:"name" short:"A"`
}

// Execute generates the supporting files file.
func (s *Support) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

// apply options.
func (s *Support) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

// generate support source.
func (s *Support) generate(opts *generator.GenOpts) error { _ = "STUB: not implemented"; return nil }

// log after generation.
func (s Support) log(_ string) { _ = "STUB: not implemented"; return }
