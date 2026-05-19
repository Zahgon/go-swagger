// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import (
	"github.com/go-swagger/go-swagger/generator"
)

type serverOptions struct {
	ServerPackage         string `default:"restapi" description:"the package to save the server specific code"                                               long:"server-package"         short:"s"`
	MainTarget            string `default:""        description:"the location of the generated main. Defaults to cmd/{name}-server"                          long:"main-package"           short:""`
	ImplementationPackage string `default:""        description:"the location of the backend implementation of the server, which will be autowired with api" long:"implementation-package" short:""`
}

func (cs serverOptions) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

// Server the command to generate an entire server application.
type Server struct {
	WithShared
	WithModels
	WithOperations

	serverOptions
	schemeOptions
	mediaOptions

	SkipModels             bool   `description:"no models will be generated when this flag is specified"           long:"skip-models"`
	SkipOperations         bool   `description:"no operations will be generated when this flag is specified"       long:"skip-operations"`
	SkipSupport            bool   `description:"no supporting files will be generated when this flag is specified" long:"skip-support"`
	ExcludeMain            bool   `description:"exclude main function, so just generate the library"               long:"exclude-main"`
	ExcludeSpec            bool   `description:"don't embed the swagger specification"                             long:"exclude-spec"`
	FlagStrategy           string `choice:"go-flags"                                                               choice:"pflag"                 choice:"flag"    default:"go-flags"                                      description:"the strategy to provide flags for the server" long:"flag-strategy"`
	CompatibilityMode      string `choice:"modern"                                                                 choice:"intermediate"          default:"modern" description:"the compatibility mode for the tls server" long:"compatibility-mode"`
	RegenerateConfigureAPI bool   `description:"Force regeneration of configureapi.go"                             long:"regenerate-configureapi"`

	Name string `description:"the name of the application, defaults to a mangled value of info.title" long:"name" short:"A"`
	// TODO(fredbi): CmdName string `long:"cmd-name" short:"A" description:"the name of the server command, when main is generated (defaults to {name}-server)"`

	// deprecated flags
	WithContext bool `description:"handlers get a context as first arg (deprecated)" long:"with-context"`
}

// Execute runs this command.
func (s *Server) Execute(_ []string) error { _ = "STUB: not implemented"; return nil }

// apply options.
func (s *Server) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

func (s *Server) generate(opts *generator.GenOpts) error { _ = "STUB: not implemented"; return nil }

func (s Server) log(_ string) { _ = "STUB: not implemented"; return }
