// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generate

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/spf13/viper"

	"github.com/go-openapi/analysis"

	"github.com/go-swagger/go-swagger/generator"
)

const (
	verboseFlag   = "verbose"
	noverboseFlag = "noverbose"
	minimalFlag   = "minimal"
	fullFlag      = "full"
)

// FlattenCmdOptions determines options to the flatten spec preprocessing.
type FlattenCmdOptions struct {
	WithExpand          bool     `description:"expands all $ref's in spec prior to generation (shorthand to --with-flatten=expand)" group:"shared" long:"with-expand"`
	WithFlatten         []string `choice:"minimal"                                                                                  choice:"full"  choice:"expand"              choice:"verbose" choice:"noverbose" choice:"remove-unused" choice:"keep-names" default:"minimal" default:"verbose" description:"flattens all $ref's in spec prior to generation" group:"shared" long:"with-flatten"`
	WithCustomFormatter bool     `description:"use faster custom contributed go import processing instead of the standard one"      group:"shared" long:"with-custom-formatter"`
}

// SetFlattenOptions builds flatten options from command line args.
func (f *FlattenCmdOptions) SetFlattenOptions(dflt *analysis.FlattenOpts) (res *analysis.FlattenOpts) {
	_ = "STUB: not implemented"
	return nil
}

// verbose flag takes precedence

// minimal flag takes precedence

// expand flag takes precedence

type sharedCommand interface {
	apply(options *generator.GenOpts)
	getConfigFile() string
	generate(options *generator.GenOpts) error
	log(command string)
}

type schemeOptions struct {
	Principal     string `description:"the model to use for the security principal" long:"principal"                              short:"P"`
	DefaultScheme string `default:"http"                                            description:"the default scheme for this API" long:"default-scheme"`

	PrincipalIface bool `description:"the security principal provided is an interface, not a struct" long:"principal-is-interface"`
}

func (so schemeOptions) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

type mediaOptions struct {
	DefaultProduces string `default:"application/json" description:"the default mime type that API operations produce" long:"default-produces"`
	DefaultConsumes string `default:"application/json" description:"the default mime type that API operations consume" long:"default-consumes"`
}

func (m mediaOptions) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

// WithShared adds the shared options group.
type WithShared struct {
	Shared sharedOptions `group:"Options common to all code generation commands"`
}

func (w WithShared) getConfigFile() string { _ = "STUB: not implemented"; return "" }

type sharedOptionsCommon struct {
	FlattenCmdOptions

	Spec                  flags.Filename `description:"the spec file to use (default swagger.{json,yml,yaml})"                             group:"shared"                                            long:"spec"                    short:"f"`
	Target                flags.Filename `default:"./"                                                                                     description:"the base directory for generating the files" group:"shared"                 long:"target"   short:"t"`
	Template              string         `choice:"stratoscale"                                                                             description:"load contributed templates"                  group:"shared"                 long:"template"`
	TemplateDir           flags.Filename `description:"alternative template override directory"                                            group:"shared"                                            long:"template-dir"            short:"T"`
	ConfigFile            flags.Filename `description:"configuration file to use for overriding template options"                          group:"shared"                                            long:"config-file"             short:"C"`
	CopyrightFile         flags.Filename `description:"copyright file used to add copyright header"                                        group:"shared"                                            long:"copyright-file"          short:"r"`
	AdditionalInitialisms []string       `description:"consecutive capitals that should be considered intialisms"                          group:"shared"                                            long:"additional-initialism"`
	AllowTemplateOverride bool           `description:"allows overriding protected templates"                                              group:"shared"                                            long:"allow-template-override"`
	SkipValidation        bool           `description:"skips validation of spec prior to generation"                                       group:"shared"                                            long:"skip-validation"`
	DumpData              bool           `description:"when present dumps the json for the template generator instead of generating files" group:"shared"                                            long:"dump-data"`
	StrictResponders      bool           `description:"Use strict type for the handler return value"                                       long:"strict-responders"`
	ReturnErrors          bool           `description:"handlers explicitly return an error as the second value"                            group:"shared"                                            long:"return-errors"           short:"e"`
}

func (s sharedOptionsCommon) apply(opts *generator.GenOpts) { _ = "STUB: not implemented"; return }

//nolint:staticcheck // tracked for migration to mangling.WithAdditionalInitialisms

func setCopyright(copyrightFile string) (string, error) {
	_ = "STUB: not implemented"
	// read the Copyright from file path in opts
	return "", nil
}

func createSwagger(s sharedCommand) error { _ = "STUB: not implemented"; return nil }

// process explicit config file argument

// viper config Debug

// TODO(fredbi): we should try and remove the need to work with relative paths,
// as this causes unnecessary constraints on os'es that support multiple drives
// (i.e. not single root like on unix), for example Windows.

func readConfig(filename string) (*viper.Viper, error) { _ = "STUB: not implemented"; return nil, nil }

func configureOptsFromConfig(cfg *viper.Viper, opts *generator.GenOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func setDebug(cfg *viper.Viper) { _ = "STUB: not implemented"; return }

// viper config debug
