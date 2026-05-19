// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package gentest

import (
	"io"
	"strings"
	"sync"
	"testing"
	"time"
)

var lockLogger sync.Mutex

// DiscardOutput discards the standard logger and returns
// a rollback function.
//
// Typical usage:
//
//	defer gentest.DiscardOutput()()
func DiscardOutput() func() { _ = "STUB: not implemented"; return nil }

// CaptureOutput captures the standard logger to the passed writer
// and returns a rollback function.
// Typical usage:
//
//	var buf bytes.Buffer
//	defer gentest.CaptureOutput(&buf)()
func CaptureOutput(w io.Writer) func() { _ = "STUB: not implemented"; return nil }

func setOutput(w io.Writer) func() { _ = "STUB: not implemented"; return nil }

// discards log output then sends a function to set it back to its original value

const minute = 60 * time.Second

// GoExecInDir executes a go commands from a target current directory.
//
// It returns a test runner func(*testing.T).
//
// Typical usage:
//
//	t.Run("should execute mycommand", gentest.GoExecInDir(folder, args))
func GoExecInDir(target string, args ...string) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

func ExecInDir(target string, command string, args ...string) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

var sanitizer = strings.NewReplacer(
	"(", "-",
	")", "-",
	".", "-",
	"_", "-",
	"\\", "/",
	":", "-",
	" ", "-",
)

func SanitizeGoModPath(pth string) string { _ = "STUB: not implemented"; return "" }

type GoModOption func(o *goModOptions)

type goModOptions struct {
	moduleName string
}

func WithGoModuleName(name string) GoModOption { _ = "STUB: not implemented"; return *new(GoModOption) }

func GoModInit(pth string, opts ...GoModOption) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // "tainted" args exec is actually okay

func GoModTidy(pth string) func(*testing.T) { _ = "STUB: not implemented"; return nil }

func GoModReplace(pth, src, dst string) func(*testing.T) { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G204: is okay for tests

func GoBuild(pth string) func(*testing.T) { _ = "STUB: not implemented"; return nil }
