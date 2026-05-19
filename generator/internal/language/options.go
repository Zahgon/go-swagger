// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package language provides the language-specific options used by the
// go-swagger code generator. The primary type is [Options], which describes
// formatting, naming, and import resolution rules for a target language.
package language

import (
	"golang.org/x/tools/imports"
)

// DefaultIndent is the default tab width used for Go source formatting.
const DefaultIndent = 2

// FormatterFunc is a function that processes go code to reformat it, e.g. [golang.org/x/tools/imports.Process]).
//
// Formatting options allow for injecting a custom formatter for the generated code. See [WithCustomFormatter].
type FormatterFunc func(filename string, src []byte, opts ...FormatOption) ([]byte, error)

// MangleFunc is a function that transforms a name string.
type MangleFunc func(string) string

// FormatOption allows for more flexible code formatting settings.
type FormatOption func(*FormatOpts)

// FormatOpts holds options for code formatting.
type FormatOpts struct {
	imports.Options

	LocalPrefixes []string
}

// WithFormatLocalPrefixes adds local prefixes to group imports.
func WithFormatLocalPrefixes(prefixes ...string) FormatOption {
	_ = "STUB: not implemented"
	return *new(FormatOption)
}

// WithFormatOnly tells the formatter to skip imports processing.
func WithFormatOnly(enabled bool) FormatOption {
	_ = "STUB: not implemented"
	return *new(FormatOption)
}

// DefaultFormatOpts is the default set of formatting options.
var DefaultFormatOpts = FormatOpts{
	Options: imports.Options{
		TabIndent: true,
		TabWidth:  DefaultIndent,
		Fragment:  true,
		Comments:  true,
	},
	LocalPrefixes: []string{"github.com/go-openapi"},
}

// FormatOptsWithDefault applies the given options on top of [DefaultFormatOpts].
func FormatOptsWithDefault(opts []FormatOption) FormatOpts {
	_ = "STUB: not implemented"
	return *new(FormatOpts)
}

// Options describes a target language to the code generator.
type Options struct {
	ReservedWords        []string
	BaseImportFunc       MangleFunc                     `json:"-"`
	ImportsFunc          func(map[string]string) string `json:"-"`
	ArrayInitializerFunc func(any) (string, error)      `json:"-"`
	FormatOnly           bool
	reservedWordsSet     map[string]struct{}
	initialized          bool
	formatFunc           FormatterFunc
	fileNameFunc         MangleFunc // language specific source file naming rules
	dirNameFunc          MangleFunc // language specific directory naming rules
}

// SetFormatFunc sets the formatting function for this language.
func (l *Options) SetFormatFunc(fn FormatterFunc) {
	_ = "STUB: not implemented"

	// Init the language option.
	return
}

func (l *Options) Init() { _ = "STUB: not implemented"; return }

// MangleName makes sure a reserved word gets a safe name.
func (l *Options) MangleName(name, suffix string) string { _ = "STUB: not implemented"; return "" }

//nolint:staticcheck // tracked for migration to mangling.NameMangler

// MangleVarName makes sure a reserved word gets a safe name.
func (l *Options) MangleVarName(name string) string { _ = "STUB: not implemented"; return "" }

//nolint:staticcheck // tracked for migration to mangling.NameMangler

// MangleFileName makes sure a file name gets a safe name.
func (l *Options) MangleFileName(name string) string { _ = "STUB: not implemented"; return "" }

//nolint:staticcheck // tracked for migration to mangling.NameMangler

// ManglePackageName makes sure a package gets a safe name.
// In case of a file system path (e.g. name contains "/" or "\" on Windows), this return only the last element.
func (l *Options) ManglePackageName(name, suffix string) string {
	_ = "STUB: not implemented"
	return ""
}

// preserve path
// drop path
//nolint:staticcheck // tracked for migration to mangling.NameMangler

// ManglePackagePath makes sure a full package path gets a safe name.
// Only the last part of the path is altered.
func (l *Options) ManglePackagePath(name string, suffix string) string {
	_ = "STUB: not implemented"
	return ""
}

// preserve path

// FormatContent formats a file with a language specific formatter.
func (l *Options) FormatContent(name string, content []byte, opts ...FormatOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unformatted content

// Imports generates the code to import some external packages, possibly aliased.
func (l *Options) Imports(imports map[string]string) string { _ = "STUB: not implemented"; return "" }

// ArrayInitializer builds a literal array.
func (l *Options) ArrayInitializer(data any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// BaseImport figures out the base path to generate import statements.
func (l *Options) BaseImport(tgt string) string { _ = "STUB: not implemented"; return "" }

// importAlias extracts the last path component from a package import path.
func importAlias(pkg string) string { _ = "STUB: not implemented"; return "" }
