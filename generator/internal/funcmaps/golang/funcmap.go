// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package golang provides the Go-specific template function map used by the
// go-swagger code generator. Functions defined here are pure utilities with
// no dependency on the generator's own types (GenSchema, GenOperation, etc.).
package golang

import (
	"strings"
	"text/template"
)

// FuncMap returns a template.FuncMap containing all Go-specific template
// functions that are independent of generator types. Callers typically
// merge additional entries (e.g. LanguageOpts-dependent or type-dependent
// functions) on top.
func FuncMap() template.FuncMap { _ = "STUB: not implemented"; return *new(template.FuncMap) }

//nolint:staticcheck // tracked for migration to mangling.NameMangler
//nolint:staticcheck // tracked for migration to mangling.NameMangler
//nolint:staticcheck // tracked for migration to mangling.NameMangler

// Pascalize converts a name to Go PascalCase, handling special prefix characters.
func Pascalize(arg string) string { _ = "STUB: not implemented"; return "" }

//nolint:staticcheck // tracked for migration to mangling.NameMangler

// PrefixForName returns a human-readable prefix for names starting with
// special characters. It is used as [swag.GoNamePrefixFunc].
func PrefixForName(arg string) string { _ = "STUB: not implemented"; return "" }

func replaceSpecialChar(in rune) string { _ = "STUB: not implemented"; return "" }

func cleanupEnumVariant(in string) string { _ = "STUB: not implemented"; return "" }

// AsJSON marshals data to a compact JSON string.
func AsJSON(data any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// AsPrettyJSON marshals data to an indented JSON string.
func AsPrettyJSON(data any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func pluralizeFirstWord(arg string) string { _ = "STUB: not implemented"; return "" }

// DropPackage returns the last component of a dot-separated name.
func DropPackage(str string) string { _ = "STUB: not implemented"; return "" }

// ContainsPkgStr returns true if str contains a package qualifier (e.g. "model.MyType").
func ContainsPkgStr(str string) bool { _ = "STUB: not implemented"; return false }

func padSurround(entry, padWith string, i, ln int) string { _ = "STUB: not implemented"; return "" }

func padComment(str string, pads ...string) string { _ = "STUB: not implemented"; return "" }

func blockComment(str string) string { _ = "STUB: not implemented"; return "" }

func dict(values ...any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // bounds checked by modulo guard above

func isInteger(arg any) bool { _ = "STUB: not implemented"; return false }

func httpStatus(code int) string { _ = "STUB: not implemented"; return "" }

func gt0(in *int64) bool { _ = "STUB: not implemented"; return false }

const (
	mdNewLine      = "</br>"
	mimeParamParts = 2
)

var (
	mdNewLineReplacer = strings.NewReplacer("\r\n", mdNewLine, "\n", mdNewLine, "\r", mdNewLine)
	interfaceReplacer = strings.NewReplacer("interface {}", "any")
)

func markdownBlock(in string) string { _ = "STUB: not implemented"; return "" }

// MediaMime extracts the MIME type from a media type string, stripping
// any parameters after the first semicolon.
func MediaMime(orig string) string { _ = "STUB: not implemented"; return "" }

// MediaGoName converts a MIME media type string to a Go-style PascalCase name.
func MediaGoName(media string) string { _ = "STUB: not implemented"; return "" }
