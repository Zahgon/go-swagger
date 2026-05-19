// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

// GenerateClient generates a client library for a swagger spec document.
func GenerateClient(name string, modelNames, operationIDs []string, opts *GenOpts) error {
	_ = "STUB: not implemented"
	return nil
}

type clientGenerator struct {
	appGenerator
}

func (c *clientGenerator) Generate() error { _ = "STUB: not implemented"; return nil }
