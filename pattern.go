package gitignore

import (
	"github.com/danwakefield/fnmatch"
	"regexp"
	"strings"
)

type Pattern interface {
	Match(path string, isDir bool) MatchResult
}

type ptrn struct {
	domain    string
	pattern   string
	re        *regexp.Regexp
	inclusion bool
	dirOnly   bool
	isFnMatch bool
}

func ParsePattern(pattern, domain string) (Pattern, error) {
	if !strings.HasPrefix(domain, "/") {
		domain = "/" + domain
	}

	p := ptrn{domain: domain}
	if strings.HasPrefix(pattern, "!") {
		p.inclusion = true
		pattern = pattern[1:]
	}

	if !strings.HasSuffix(pattern, "\\ ") {
		pattern = strings.TrimRight(pattern, " ")
	}

	if strings.HasSuffix(pattern, "/") {
		p.dirOnly = true
		pattern = pattern[:len(pattern)-1]
	} else {
		if re, err := regexp.Compile(p.pattern); err != nil {
			return nil, err
		} else {
			p.re = re
		}
	}

	if strings.Contains(pattern, "/") {
		p.isFnMatch = true
	}

	if strings.HasPrefix(pattern, "/") {
		pattern = domain + pattern
	}

	p.pattern = pattern
	return &p, nil
}

func (p *ptrn) Match(path string, isDir bool) MatchResult {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !strings.HasPrefix(path, p.domain) || (p.dirOnly && !isDir) {
		return NoMatch
	}
	if p.isFnMatch {
		if !fnmatch.Match(p.pattern, path, fnmatch.FNM_PATHNAME) {
			return NoMatch
		}
	} else {
		if !p.re.Match([]byte(path)) {
			return NoMatch
		}
	}

	if p.inclusion {
		return Include
	} else {
		return Exclude
	}
}
