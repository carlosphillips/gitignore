package gitignore

import (
	"path/filepath"
	"strings"
)

type Pattern interface {
	Match(path []string, isDir bool) MatchResult
}

type ptrn struct {
	domain    []string
	pattern   []string
	inclusion bool
	dirOnly   bool
	isGlob    bool
}

func ParsePattern(pattern string, domain []string) Pattern {
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
	}

	if strings.Contains(pattern, "/") {
		p.isGlob = true
	}

	p.pattern = strings.Split(pattern, "/")
	return &p
}

func (p *ptrn) Match(path []string, isDir bool) MatchResult {
	if len(path) <= len(p.domain) {
		return NoMatch
	}
	for i, e := range p.domain {
		if path[i] != e {
			return NoMatch
		}
	}

	path = path[len(p.domain):]
	if p.isGlob && !p.globMatch(path, isDir) {
		return NoMatch
	} else if !p.isGlob && !p.simpleNameMatch(path, isDir) {
		return NoMatch
	}

	if p.inclusion {
		return Include
	} else {
		return Exclude
	}
}

func (p *ptrn) simpleNameMatch(path []string, isDir bool) bool {
	for i, name := range path {
		if match, err := filepath.Match(p.pattern[0], name); err != nil {
			return false
		} else if !match {
			continue
		}
		if p.dirOnly && !isDir && i == len(path)-1 {
			return false
		}
		return true
	}
	return false
}

func (p *ptrn) globMatch(path []string, isDir bool) bool {
	matched := false
	canTraverse := false
	for i, pattern := range p.pattern {
		if pattern == "" {
			canTraverse = false
			continue
		}
		if pattern == "**" {
			if i == len(p.pattern)-1 {
				break
			}
			canTraverse = true
			continue
		}
		if strings.Contains(pattern, "**") {
			return false
		}
		if len(path) == 0 {
			return false
		}
		if canTraverse {
			canTraverse = false
			for len(path) > 0 {
				e := path[0]
				path = path[1:]
				if match, err := filepath.Match(pattern, e); err != nil {
					return false
				} else if match {
					matched = true
					break
				}
			}
		} else {
			if match, err := filepath.Match(pattern, path[0]); err != nil || !match {
				return false
			}
			matched = true
			path = path[1:]
		}
	}
	if matched && p.dirOnly && !isDir && len(path) == 0 {
		matched = false
	}
	return matched
}
