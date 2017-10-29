// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package gitignore

import (
	"strings"
)

// Dir defines a directory structure from which files can be loaded and sub-directories traversed.
type Dir interface {
	// Path returns the path to this directory
	Path() []string
	// ReadFile reads the content of a file in this directory if file exists.
	ReadFile(name string) ([]byte, error)
	// Subdirs lits sub-directories.
	Subdirs() ([]Dir, error)
}

// ReadPatterns reads gitignore patterns recursively traversing through the directory
// structure. The result is in the ascending order of priority (last higher).
func ReadPatterns(dir Dir) (patterns []Pattern, err error) {
	if data, err := dir.ReadFile(".gitignore"); err == nil {
		for _, s := range strings.Split(string(data), "\n") {
			if !strings.HasPrefix(s, "#") && len(strings.TrimSpace(s)) > 0 {
				patterns = append(patterns, ParsePattern(s, dir.Path()))
			}
		}
	}

	var subdirs []Dir
	subdirs, err = dir.Subdirs()
	if err != nil {
		return
	}
	for _, subdir := range subdirs {
		if subdir.Path()[len(subdir.Path())-1] != ".git" {
			var subpatterns []Pattern
			subpatterns, err = ReadPatterns(subdir)
			if err != nil {
				return
			}
			if len(subpatterns) > 0 {
				patterns = append(patterns, subpatterns...)
			}
		}
	}

	return
}
