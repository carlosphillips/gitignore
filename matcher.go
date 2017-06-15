package gitignore

type MatchResult int

const (
	NoMatch MatchResult = iota
	Exclude
	Include
)

type Matcher interface {
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

// Match matches patterns in the order of priorities. As soon as an inclusion or
// exclusion is found, not further matching is performed.
func (m *matcher) Match(path []string, isDir bool) bool {
	n := len(m.patterns)
	for i := n - 1; i >= 0; i-- {
		if match := m.patterns[i].Match(path, isDir); match > NoMatch {
			return match == Exclude
		}
	}
	return false
}
