// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

// GenerateServer generates a server application.
func GenerateServer(name string, modelNames, operationIDs []string, opts *GenOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateSupport generates the supporting files for an API.
func GenerateSupport(name string, modelNames, operationIDs []string, opts *GenOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateMarkdown documentation for a swagger specification.
func GenerateMarkdown(output string, modelNames, operationIDs []string, opts *GenOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func newAppGenerator(name string, modelNames, operationIDs []string, opts *GenOpts) (*appGenerator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default target for the generated main

type appGenerator struct {
	Name              string
	Receiver          string
	SpecDoc           *loads.Document
	Analyzed          *analysis.Spec
	Package           string
	APIPackage        string
	ModelsPackage     string
	ServerPackage     string
	ClientPackage     string
	OperationsPackage string
	MainPackage       string
	Principal         string
	Models            map[string]spec.Schema
	Operations        map[string]opRef
	Target            string
	DumpData          bool
	DefaultScheme     string
	DefaultProduces   string
	DefaultConsumes   string
	GenOpts           *GenOpts
}

func (a *appGenerator) Generate() error { _ = "STUB: not implemented"; return nil }

// NOTE: relative to previous implem with chan.
// IPC removed concurrent execution because of the FuncMap that is being shared
// templates are now lazy loaded so there is concurrent map access I can't guard

// optional OperationGroups templates generation

func (a *appGenerator) GenerateSupport(ap *GenApp) error { _ = "STUB: not implemented"; return nil }

// allows for calling GenerateSupport standalone

// no need to add this import when there is no CLI
// add client import for cli generation

func (a *appGenerator) GenerateMarkdown() error { _ = "STUB: not implemented"; return nil }

func (a *appGenerator) makeSecuritySchemes() GenSecuritySchemes {
	_ = "STUB: not implemented"
	return *new(GenSecuritySchemes)
}

//nolint:gocognit,gocyclo,cyclop,maintidx // TODO(fredbi): refactor
func (a *appGenerator) makeCodegenApp() (GenApp, error) {
	_ = "STUB: not implemented"
	return *new(GenApp), nil
}

// we don't want to inject this import for clients

// Copy model imports to operation imports
// TODO(fredbi): mangle model pkg aliases

// defaults to main operations package

// operation filtered according to CLI params

// ordered tags for this operation, possibly filtered by CLI params

// check for possible conflicts that requires import aliasing

// we don't want import to shadow the current package

// was already imported with a different target

// trim duplicate extra schemas within the same package

// generating extra options to switch media type in client

// top level securityRequirements

// generateReadableSpec makes swagger json spec as a string instead of bytes
// the only character that needs to be escaped is '`' symbol, since it cannot be escaped in the GO string
// that is quoted as `string data`. The function doesn't care about the beginning or the ending of the
// string it escapes since all data that needs to be escaped is always in the middle of the swagger spec.
func generateReadableSpec(spec []byte) string { _ = "STUB: not implemented"; return "" }

func trimExternalDoc(in *spec.ExternalDocumentation) *spec.ExternalDocumentation {
	_ = "STUB: not implemented"
	return nil
}

func trimInfo(in *spec.Info) *spec.Info { _ = "STUB: not implemented"; return nil }

func trimTags(in []spec.Tag) []spec.Tag { _ = "STUB: not implemented"; return nil }
