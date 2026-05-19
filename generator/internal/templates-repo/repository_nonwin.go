// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package templatesrepo

// LoadPlugin will load the named plugin and inject its functions into the funcMap
//
// The plugin must implement a function matching the signature:
// `func AddFuncs(f template.FuncMap)`
// which can add any number of functions to the template repository funcMap.
// Any existing sprig or go-swagger templates with the same name will be overridden.
func (t *Repository) LoadPlugin(pluginPath string) error { _ = "STUB: not implemented"; return nil }
