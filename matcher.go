package gitignore

type MatchResult int

const (
	NoMatch MatchResult = iota
	Exclude
	Include
)

type Matcher interface {
	Match(path string, isDir bool) MatchResult
}

func NewMatcher(patterns []Pattern) Matcher {
	return &matcher{patterns}
}

type matcher struct {
	patterns []Pattern
}

func (m *matcher) Match(path string, isDir bool) MatchResult {
	var res MatchResult
	for _, p := range m.patterns {
		if pres := p.Match(path, isDir); pres > NoMatch {
			res = pres
		}
	}
	return res
}
