// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package cmdtest provides test utilities
// to assert the output of CLI commands
package cmdtest

import (
	"io"
	"testing"
)

// AssertReadersContent compares the contents from io.Readers, optionally stripping blanks.
func AssertReadersContent(t testing.TB, noBlanks bool, expected, actual io.Reader) bool {
	_ = "STUB: not implemented"
	return false
}

// CatchStdOut captures the standard output from a runnable function.
// You shouln't run this in parallel.
func CatchStdOut(t testing.TB, runnable func() error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// need to close here, otherwise ReadAll never gets "EOF".
