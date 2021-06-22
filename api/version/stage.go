// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

// Stage represents the the different version stages that
// can be defined in a migration such as a Patch,
// Minor or Major.
type Stage string

const (
	// Major signals backward-incompatible public API changes.
	// This release carries no guarantee that it will be
	// backward compatible with preceding major versions.
	Major = "major"
	// Minor signals backward-compatible public API changes.
	// This release guarantees backward compatibility and
	// stability.
	Minor = "minor"
	// Patch signals changes that don't affect the module's
	// public API or its dependencies. This release
	// guarantees backward compatibility and
	// stability.
	Patch = "patch"
)
