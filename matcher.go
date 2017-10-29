// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

// gitignore package implements scanning for and parsing of gitignore files in a directory
// structure and matching parsed patterns to filesystem paths. All patterns listed in the
// original gitignore specification are supported.
package gitignore

// Matcher defines a global multi-pattern matcher for gitignore patterns
type Matcher interface {
	// Match matches patterns in the order of priorities. As soon as an inclusion or
	// exclusion is found, not further matching is performed.
	Match(path []string, isDir bool) bool
}

// NewMatcher constructs a new global matcher. Patterns must be given in the order of
// increasing priority. That is most generic settings files first, then the content of
// the repo .gitignore, then content of .gitignore down the path or the repo and then
// the content command line arguments.
func NewMatcher(patterns []Pattern) Matcher {
	return &matcher{patterns}
}

type matcher struct {
	patterns []Pattern
}

func (m *matcher) Match(path []string, isDir bool) bool {
	n := len(m.patterns)
	for i := n - 1; i >= 0; i-- {
		if match := m.patterns[i].Match(path, isDir); match > NoMatch {
			return match == Exclude
		}
	}
	return false
}
