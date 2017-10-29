// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package gitignore_test

import (
	"github.com/teris-io/gitignore"
	"testing"
)

func TestMatcher_Match(t *testing.T) {
	patterns := []gitignore.Pattern{
		gitignore.ParsePattern("**/middle/v[uo]l?ano", nil),
		gitignore.ParsePattern("!volcano", nil),
	}
	matcher := gitignore.NewMatcher(patterns)
	if !matcher.Match([]string{"head", "middle", "vulkano"}, false) {
		t.Errorf("expected a match, found mismatch")
	}
	if matcher.Match([]string{"head", "middle", "volcano"}, false) {
		t.Errorf("expected a mismatch, found a match")
	}
}
