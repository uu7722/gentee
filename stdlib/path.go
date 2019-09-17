// Copyright 2019 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package stdlib

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gentee/gentee/core"
)

// InitPath appends stdlib filepath functions to the virtual machine
func InitPath(ws *core.Workspace) {
	for _, item := range []interface{}{
		core.Link{AbsPath, 97<<16 | core.EMBED},    // AbsPath(str) str
		core.Link{BaseName, 98<<16 | core.EMBED},   // BaseName(str) str
		core.Link{Dir, 99<<16 | core.EMBED},        // Dir(str) str
		core.Link{Ext, 100<<16 | core.EMBED},       // Ext(str) str
		core.Link{JoinPath, 101<<16 | core.EMBED},  // JoinPath(str pars...) str
		core.Link{MatchPath, 102<<16 | core.EMBED}, // MatchPath(str, str) bool
	} {
		ws.StdLib().NewEmbed(item)
	}
}

// AbsPath returns an absolute representation of path.
func AbsPath(fname string) (string, error) {
	return filepath.Abs(fname)
}

// BaseName returns the last element of path.
func BaseName(fname string) string {
	return filepath.Base(fname)
}

// Dir returns all but the last element of path.
func Dir(fname string) string {
	return filepath.Dir(fname)
}

// Ext returns the file name extension used by path.
func Ext(fname string) string {
	return strings.TrimLeft(filepath.Ext(fname), `.`)
}

// JoinPath joins any number of path elements into a single path.
func JoinPath(pars ...interface{}) string {
	names := make([]string, len(pars))
	for i, name := range pars {
		names[i] = fmt.Sprint(name)
	}
	return filepath.Join(names...)
}

// MatchPath reports whether name matches the specified file name pattern.
func MatchPath(pattern, fname string) (bool, error) {
	return filepath.Match(pattern, fname)
}
