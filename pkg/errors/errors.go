package errors

import (
	"errors"

	"golang.org/x/xerrors"
)

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
//
//nolint:gochecknoglobals
var New = errors.New

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
//
//nolint:gochecknoglobals
var Errorf = xerrors.Errorf

// Is reports whether any error in err's chain matches target.
//
//nolint:gochecknoglobals
var Is = errors.Is

// As finds the first error in err's chain that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
//
//nolint:gochecknoglobals
var As = errors.As
