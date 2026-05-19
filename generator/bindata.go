// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"embed"
)

//go:embed templates
var _bindata embed.FS

// AssetNames returns the names of the assets.
func AssetNames() []string { _ = "STUB: not implemented"; return nil }

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte { _ = "STUB: not implemented"; return nil }
