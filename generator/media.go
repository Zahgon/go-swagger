// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"regexp"
)

const jsonSerializer = "json"

var mediaTypeNames = map[*regexp.Regexp]string{
	regexp.MustCompile("application/.*json"):                jsonSerializer,
	regexp.MustCompile("application/.*yaml"):                "yaml",
	regexp.MustCompile("application/.*protobuf"):            "protobuf",
	regexp.MustCompile("application/.*capnproto"):           "capnproto",
	regexp.MustCompile("application/.*thrift"):              "thrift",
	regexp.MustCompile("(?:application|text)/.*xml"):        "xml",
	regexp.MustCompile("text/.*markdown"):                   "markdown",
	regexp.MustCompile("text/.*html"):                       "html",
	regexp.MustCompile("text/.*csv"):                        "csv",
	regexp.MustCompile("text/.*tsv"):                        "tsv",
	regexp.MustCompile("text/.*javascript"):                 "js",
	regexp.MustCompile("text/.*css"):                        "css",
	regexp.MustCompile("text/.*plain"):                      "txt",
	regexp.MustCompile("application/.*octet-stream"):        "bin",
	regexp.MustCompile("application/.*tar"):                 "tar",
	regexp.MustCompile("application/.*gzip"):                "gzip",
	regexp.MustCompile("application/.*gz"):                  "gzip",
	regexp.MustCompile("application/.*raw-stream"):          "bin",
	regexp.MustCompile("application/x-www-form-urlencoded"): "urlform",
	regexp.MustCompile("application/javascript"):            "txt",
	regexp.MustCompile("multipart/form-data"):               "multipartform",
	regexp.MustCompile("image/.*"):                          "bin",
	regexp.MustCompile("audio/.*"):                          "bin",
	regexp.MustCompile("application/pdf"):                   "bin",
}

var knownProducers = map[string]string{
	jsonSerializer:  "runtime.JSONProducer()",
	"yaml":          "yamlpc.YAMLProducer()",
	"xml":           "runtime.XMLProducer()",
	"txt":           "runtime.TextProducer()",
	"bin":           "runtime.ByteStreamProducer()",
	"csv":           "runtime.CSVProducer()",
	"urlform":       "runtime.DiscardProducer",
	"multipartform": "runtime.DiscardProducer",
}

var knownConsumers = map[string]string{
	jsonSerializer:  "runtime.JSONConsumer()",
	"yaml":          "yamlpc.YAMLConsumer()",
	"xml":           "runtime.XMLConsumer()",
	"txt":           "runtime.TextConsumer()",
	"bin":           "runtime.ByteStreamConsumer()",
	"csv":           "runtime.CSVConsumer()",
	"urlform":       "runtime.DiscardConsumer",
	"multipartform": "runtime.DiscardConsumer",
}

func wellKnownMime(tn string) (string, bool) { _ = "STUB: not implemented"; return "", false }

const mimeParamParts = 2

func mediaParameters(orig string) string { _ = "STUB: not implemented"; return "" }

func (a *appGenerator) makeSerializers(mediaTypes []string, known func(string) (string, bool)) (GenSerGroups, bool) {
	_ = "STUB: not implemented"
	return *new(GenSerGroups), false
}

// build all required serializers

// keep this serializer named, even though its implementation is empty (cf. #1557)

// provide all known parameters (currently unused by codegen templates)

// group serializers by consumer/producer to serve several mime media types

// provides the full list of mime media types for this serializer group

func (a *appGenerator) makeConsumes() (GenSerGroups, bool) {
	_ = "STUB: not implemented"
	// builds a codegen struct from all consumes in the spec
	return *new(GenSerGroups), false
}

func (a *appGenerator) makeProduces() (GenSerGroups, bool) {
	_ = "STUB: not implemented"
	// builds a codegen struct from all produces in the spec
	return *new(GenSerGroups), false
}
