// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	golangfuncs "github.com/go-swagger/go-swagger/generator/internal/funcmaps/golang"
)

// Package-level aliases for functions that moved to the golang funcmap package
// but are still referenced from other files in the generator package.
var (
	pascalize   = golangfuncs.Pascalize
	mediaMime   = golangfuncs.MediaMime
	mediaGoName = golangfuncs.MediaGoName
	asJSON      = golangfuncs.AsJSON
)

func resolvedDocCollectionFormat(cf string, child *GenItems) string {
	_ = "STUB: not implemented"
	return ""
}

func resolvedDocType(tn, ft string, child *GenItems) string { _ = "STUB: not implemented"; return "" }

func resolvedDocSchemaType(tn, ft string, child *GenSchema) string {
	_ = "STUB: not implemented"
	return ""
}

func resolvedDocElemType(tn, ft string, schema *resolvedType) string {
	_ = "STUB: not implemented"
	return ""
}

func errorPath(in any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// unchanged Path if called with other types
