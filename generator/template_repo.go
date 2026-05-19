// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"text/template"

	templatesrepo "github.com/go-swagger/go-swagger/generator/internal/templates-repo"
)

var (
	assets             map[string][]byte
	protectedTemplates map[string]bool

	// FuncMapFunc yields a map with all functions for templates.
	FuncMapFunc func(*LanguageOpts) template.FuncMap

	templates *templatesrepo.Repository

	docFormat map[string]string

	errInternal = errors.New("internal error detected in templates")
)

// embeddedAssets adapts the package-level AssetNames/MustAsset functions
// to the [templatesrepo.AssetProvider] interface.
type embeddedAssets struct{}

func (embeddedAssets) AssetNames() []string         { _ = "STUB: not implemented"; return nil }
func (embeddedAssets) MustAsset(name string) []byte { _ = "STUB: not implemented"; return nil }

func initTemplateRepo() { _ = "STUB: not implemented"; return }

// this makes the ToGoName func behave with the special
// prefixing rule above
//nolint:staticcheck // tracked for migration to mangling.WithGoNamePrefixFunc

// DefaultFuncMap yields a map with default functions for use in the templates.
// These are available in every template.
func DefaultFuncMap(lang *LanguageOpts) template.FuncMap {
	_ = "STUB: not implemented"
	return *new(template.FuncMap)
}

// Language-specific entries that depend on *LanguageOpts.

// Generator-type-dependent entries.

// CLI command helpers that depend on generator types.

// assert is used to inject into templates and check for inconsistent/invalid data.

func defaultAssets() map[string][]byte { _ = "STUB: not implemented"; return nil }

// schema validation templates

// schema serialization templates

// schema generation template

// simple schema generation helpers templates

// server templates

// client templates

// cli templates

func defaultProtectedTemplates() map[string]bool { _ = "STUB: not implemented"; return nil }

// validation helpers

// all serializers
