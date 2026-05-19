// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"github.com/spf13/viper"
)

// LanguageDefinition in the configuration file.
type LanguageDefinition struct {
	Layout SectionOpts `mapstructure:"layout"`
}

// ConfigureOpts for generation.
func (d *LanguageDefinition) ConfigureOpts(opts *GenOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// LanguageConfig structure that is obtained from parsing a config file.
type LanguageConfig map[string]LanguageDefinition

// ReadConfig at the specified path, when no path is specified it will look into
// the current directory and load a .swagger.{yml,json,hcl,toml,properties} file
// Returns a viper config or an error.
func ReadConfig(fpath string) (*viper.Viper, error) { _ = "STUB: not implemented"; return nil, nil }
