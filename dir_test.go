/**
 * Copyright (c) Oleg Sklyar, Silvertern, 2017. MIT license
 */
package gitignore_test

import (
	"github.com/silvertern/gitignore"
	"testing"
	"fmt"
)

type dir struct {
	path []string
	subdirError bool
}

func (d *dir) Path() []string {
	return d.path
}

func (d *dir) ReadFile(name string) ([]byte, error) {
	if name != ".gitignore" {
		return nil, fmt.Errorf("no such file")
	}
	if len(d.path) == 0 {
		return []byte("vendor/\n"), nil
	} else if d.path[len(d.path)-1] == "vendor" {
		return []byte("!github.com/\n"), nil
	}
	return nil, fmt.Errorf("no such file")
}

func (d *dir) Subdirs() ([]gitignore.Dir, error) {
	if len(d.path) == 0 {
		return []gitignore.Dir{&dir{path: append(d.path, "vendor"), subdirError: d.subdirError}, &dir{path: append(d.path, "another"), subdirError: d.subdirError}}, nil
	} else if d.path[len(d.path)-1] == "vendor" {
		return []gitignore.Dir{&dir{path: append(d.path, "github.com"), subdirError: d.subdirError}, &dir{path: append(d.path, "gopkg.in"), subdirError: d.subdirError}}, nil
	}
	if d.subdirError {
		return nil, fmt.Errorf("failed to list directories")
	}
	return nil, nil
}

func TestDir_ReadPatterns(t *testing.T) {
	patterns, err := gitignore.ReadPatterns(&dir{})
	if err != nil {
		t.Errorf("no error expected, found %v", err)
	}
	if len(patterns) != 2 {
		t.Errorf("expected 2 patterns, found %v", len(patterns))
	}
	matcher := gitignore.NewMatcher(patterns)
	if !matcher.Match([]string{"vendor"}, true) {
		t.Error("expected a match")
	}
	if !matcher.Match([]string{"vendor", "gopkg.in"}, true) {
		t.Error("expected a match")
	}
	if matcher.Match([]string{"vendor", "github.com"}, true) {
		t.Error("expected no match")
	}
}

func TestDir_ReadPatterns_error(t *testing.T) {
	_, err := gitignore.ReadPatterns(&dir{subdirError: true})
	if err == nil {
		t.Errorf("expected an error")
	} else if err.Error() != "failed to list directories" {
		t.Errorf("expecte different error message, found %v", err)
	}

}
