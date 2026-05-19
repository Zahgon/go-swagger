// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/spec"
)

type discInfo struct {
	Discriminators map[string]discor
	Discriminated  map[string]discee
}

type discor struct {
	FieldName string   `json:"fieldName"`
	GoType    string   `json:"goType"`
	JSONName  string   `json:"jsonName"`
	Children  []discee `json:"children"`
}

type discee struct {
	FieldName  string   `json:"fieldName"`
	FieldValue string   `json:"fieldValue"`
	GoType     string   `json:"goType"`
	JSONName   string   `json:"jsonName"`
	Ref        spec.Ref `json:"ref"`
	ParentRef  spec.Ref `json:"parentRef"`
}

func discriminatorInfo(doc *analysis.Spec) *discInfo { _ = "STUB: not implemented"; return nil }
