// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package templatesrepo

import (
	"sync"
	"text/template"
	"text/template/parse"
)

// AssetProvider provides access to embedded template assets.
type AssetProvider interface {
	AssetNames() []string
	MustAsset(name string) []byte
}

// Repository is the repository for the generator templates.
type Repository struct {
	files              map[string]string
	templates          map[string]*template.Template
	funcs              template.FuncMap
	protectedTemplates map[string]bool
	allowOverride      bool
	mux                sync.Mutex
}

// NewRepository creates a new template repository with the provided functions defined.
func NewRepository(funcs template.FuncMap) *Repository { _ = "STUB: not implemented"; return nil }

// SetProtectedTemplates sets the map of template names that cannot be overridden
// by user-provided templates.
func (t *Repository) SetProtectedTemplates(m map[string]bool) { _ = "STUB: not implemented"; return }

// ShallowClone a repository.
//
// Clones the maps of files and templates, so as to be able to use
// the cloned repo concurrently.
func (t *Repository) ShallowClone() *Repository { _ = "STUB: not implemented"; return nil }

// LoadDefaults loads templates from the given asset map.
func (t *Repository) LoadDefaults(assets map[string][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadDir will walk the specified path and add each .gotmpl file it finds to the repository.
func (t *Repository) LoadDir(templatePath string) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // pre-existing: template loading from user-specified directory

// Non-readable files are skipped

// Non-template files are skipped

// LoadContrib loads template from contrib directory using the given asset provider.
func (t *Repository) LoadContrib(name string, provider AssetProvider) error {
	_ = "STUB: not implemented"
	return nil
}

// MustGet a template by name, panics when fails.
func (t *Repository) MustGet(name string) *template.Template { _ = "STUB: not implemented"; return nil }

// AddFile adds a file to the repository. It will create a new template based on the filename.
// It trims the .gotmpl from the end and converts the name using swag.ToJSONName. This will strip
// directory separators and Camelcase the next letter.
// e.g validation/primitive.gotmpl will become validationPrimitive
//
// If the file contains a definition for a template that is protected the whole file will not be added.
func (t *Repository) AddFile(name, data string) error { _ = "STUB: not implemented"; return nil }

// SetAllowOverride allows setting allowOverride after the Repository was initialized.
func (t *Repository) SetAllowOverride(value bool) { _ = "STUB: not implemented"; return }

// Get will return the named template from the repository, ensuring that all dependent templates are loaded.
// It will return an error if a dependent template is not defined in the repository.
func (t *Repository) Get(name string) (*template.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DumpTemplates prints out a dump of all the defined templates, where they are defined and what their dependencies are.
func (t *Repository) DumpTemplates() { _ = "STUB: not implemented"; return }

// Funcs returns the template function map, allowing callers to add or modify functions.
func (t *Repository) Funcs() template.FuncMap {
	_ = "STUB: not implemented"
	return *new(template.FuncMap)
}

func (t *Repository) addFile(name, data string, allowOverride bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck // tracked for migration to mangling.NameMangler

// check if any protected templates are defined

// Add each defined template into the cache

func (t *Repository) flattenDependencies(templ *template.Template, dependencies map[string]bool) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func (t *Repository) addDependencies(templ *template.Template) (*template.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if we have it

// Still don't have it, return an error

// Add it to the parse tree

func findDependencies(n parse.Node) []string { _ = "STUB: not implemented"; return nil }
