// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import (
	"github.com/go-swagger/go-swagger/generator"
)

type clientOptions struct {
	ClientPackage string `default:"client" description:"the package to save the client specific code" long:"client-package" short:"c"`
}

func (co clientOptions) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

// Client the command to generate a swagger client.
type Client struct {
	WithShared
	WithModels
	WithOperations

	clientOptions
	schemeOptions
	mediaOptions

	SkipModels     bool `description:"no models will be generated when this flag is specified"     long:"skip-models"`
	SkipOperations bool `description:"no operations will be generated when this flag is specified" long:"skip-operations"`

	Name string `description:"the name of the application, defaults to a mangled value of info.title" long:"name" short:"A"`
}

// Execute runs this command.
func (c *Client) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

// apply options.
func (c Client) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

func (c *Client) generate(opts *generator.GenOpts) error { _ = "STUB: not implemented"; return nil }

func (c *Client) log(_ string) { _ = "STUB: not implemented"; return }
