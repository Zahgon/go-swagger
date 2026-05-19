// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package generator

type GenOpts struct {
	GenOptsCommon

	TemplatePlugin string
}

func (g *GenOpts) setTemplates() error { _ = "STUB: not implemented"; return nil }
