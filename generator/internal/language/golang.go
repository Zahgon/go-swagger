// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package language

import (
	"os"
	"regexp"
)

var moduleRe = regexp.MustCompile(`module[ \t]+([^\s]+)`)

// GolangOpts returns [Options] for rendering items as golang code.
func GolangOpts() *Options { _ = "STUB: not implemented"; return nil }

// this default may be overridden by [GenOptsCommon]

func defaultGoFormatFunc() FormatterFunc { _ = "STUB: not implemented"; return *new(FormatterFunc) }

// regroup these packages

func defaultGoFilenameFunc(reservedSuffixes map[string]bool) MangleFunc {
	_ = "STUB: not implemented"
	return *new(MangleFunc)
}

//nolint:staticcheck // tracked for migration to mangling.NameMangler

func defaultGoDirnameFunc() MangleFunc { _ = "STUB: not implemented"; return *new(MangleFunc) }

func defaultGoImportsFunc() func(map[string]string) string { _ = "STUB: not implemented"; return nil }

func defaultGoArrayInitializerFunc() func(any) (string, error) {
	_ = "STUB: not implemented"
	return nil
}

func defaultGoBaseImportFunc() MangleFunc { _ = "STUB: not implemented"; return *new(MangleFunc) }

// NOTE: historically this called log.Fatalln. We now panic to avoid
// pulling in generator-specific logging, while preserving the "fail hard" semantics.

// DefaultGoBaseImportErr resolves the Go import path for the given target directory.
func DefaultGoBaseImportErr(target string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func defaultGoBaseImportErr(target string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func exploreGoPath(gopath, targetAbsPath, targetAbsPathExtended string) (pth string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint:gosec // GOPATH traversal is expected

func resolveGoModFile(dir string) (*os.File, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func relPathToRelGoPath(modAbsPath, absPath string) string { _ = "STUB: not implemented"; return "" }

func tryResolveModule(baseTargetPath string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// CheckPrefixAndFetchRelativePath checks if childpath is under parentpath
// and returns the relative path if so.
func CheckPrefixAndFetchRelativePath(childpath string, parentpath string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func goOtherReservedSuffixes() map[string]bool { _ = "STUB: not implemented"; return nil }

// goos

// arch

// other reserved suffixes
