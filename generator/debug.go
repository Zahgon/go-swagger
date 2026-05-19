// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"log"
	"os"
)

var (
	// Debug when the env var DEBUG or SWAGGER_DEBUG is not empty
	// the generators will be very noisy about what they are doing.
	Debug = os.Getenv("DEBUG") != "" || os.Getenv("SWAGGER_DEBUG") != ""
	// generatorLogger is a debug logger for this package.
	generatorLogger *log.Logger
)

func debugOptions() { _ = "STUB: not implemented"; return }

// debugLog wraps log.Printf with a debug-specific logger.
func debugLogf(format string, args ...any) { _ = "STUB: not implemented"; return }

// debugLogAsJSON unmarshals its last arg as pretty JSON.
func debugLogAsJSONf(format string, args ...any) { _ = "STUB: not implemented"; return }

//nolint:errchkjson // OK: it's okay for debug

// sanitizeDebugLogArgs traverses arguments to debugLog and redacts fields
// that may contain sensitive information, such as API keys or credentials.
func sanitizeDebugLogArgs(args ...any) []any { _ = "STUB: not implemented"; return nil }

// sanitizeValue redacts sensitive information from known data structures.
// It can be expanded for more types over time as needed.
func sanitizeValue(val any) any { _ = "STUB: not implemented"; return *new(any) }

// recursively sanitize map values

// false positive: this is a bool indicator, not a sensitive value

// heuristic: redact if looks like a key/secret

// Optionally, process struct types for known sensitive fields

// fatal wraps [log.Fatal] with extra context provided in debug mode.
func fatal(v ...any) {
	_ = "STUB: not implemented"

	// fatalln wraps [log.Fatalln] with extra context provided in debug mode.
	return
}

func fatalln(v ...any) { _ = "STUB: not implemented"; return }

// traceFatalf allows to capture more context about the caller of a fatalX function.
//
// This output is not disabled when muting the [log.Logger] (e.g. when running tests).
func traceFatalf(format string, v ...any) { _ = "STUB: not implemented"; return }
