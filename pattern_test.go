package gitignore_test

import (
	"github.com/silvertern/gitignore"
	"testing"
)

func TestPattern_fnMatch_withDoubleStars(t *testing.T) {
	p, _ := gitignore.ParsePattern("**/foo", "ab/bc/de")
	if match := p.Match("/ab/bc/de/foo", false); match != gitignore.Exclude {
		t.Errorf("expected 1, found %v", match)
	}
}
