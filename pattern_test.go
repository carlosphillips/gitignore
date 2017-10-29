// Copyright (c) 2017. Oleg Sklyar & teris.io. All rights reserved.
// See the LICENSE file in the project root for licensing information.

package gitignore_test

import (
	"github.com/teris-io/gitignore"
	"testing"
)

func TestPatternSimpleMatch_inclusion(t *testing.T) {
	pattern := gitignore.ParsePattern("!vul?ano", nil)
	if res := pattern.Match([]string{"value", "vulkano", "tail"}, false); res != gitignore.Include {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternMatch_domainLonger_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("value", []string{"head", "middle", "tail"})
	if res := pattern.Match([]string{"head", "middle"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternMatch_domainSameLength_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("value", []string{"head", "middle", "tail"})
	if res := pattern.Match([]string{"head", "middle", "tail"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternMatch_domainMismatch_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("value", []string{"head", "middle", "tail"})
	if res := pattern.Match([]string{"head", "middle", "_tail_", "value"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternSimpleMatch_withDomain(t *testing.T) {
	pattern := gitignore.ParsePattern("middle/", []string{"value", "volcano"})
	if res := pattern.Match([]string{"value", "volcano", "middle", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_onlyMatchInDomain_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("volcano/", []string{"value", "volcano"})
	if res := pattern.Match([]string{"value", "volcano", "tail"}, true); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternSimpleMatch_atStart(t *testing.T) {
	pattern := gitignore.ParsePattern("value", nil)
	if res := pattern.Match([]string{"value", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_inTheMiddle(t *testing.T) {
	pattern := gitignore.ParsePattern("value", nil)
	if res := pattern.Match([]string{"head", "value", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_atEnd(t *testing.T) {
	pattern := gitignore.ParsePattern("value", nil)
	if res := pattern.Match([]string{"head", "value"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_atStart_dirWanted(t *testing.T) {
	pattern := gitignore.ParsePattern("value/", nil)
	if res := pattern.Match([]string{"value", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_inTheMiddle_dirWanted(t *testing.T) {
	pattern := gitignore.ParsePattern("value/", nil)
	if res := pattern.Match([]string{"head", "value", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_atEnd_dirWanted(t *testing.T) {
	pattern := gitignore.ParsePattern("value/", nil)
	if res := pattern.Match([]string{"head", "value"}, true); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_atEnd_dirWanted_notADir_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("value/", nil)
	if res := pattern.Match([]string{"head", "value"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternSimpleMatch_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("value", nil)
	if res := pattern.Match([]string{"head", "val", "tail"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternSimpleMatch_valueLonger_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("val", nil)
	if res := pattern.Match([]string{"head", "value", "tail"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternSimpleMatch_withAsterisk(t *testing.T) {
	pattern := gitignore.ParsePattern("v*o", nil)
	if res := pattern.Match([]string{"value", "vulkano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_withQuestionMark(t *testing.T) {
	pattern := gitignore.ParsePattern("vul?ano", nil)
	if res := pattern.Match([]string{"value", "vulkano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_magicChars(t *testing.T) {
	pattern := gitignore.ParsePattern("v[ou]l[kc]ano", nil)
	if res := pattern.Match([]string{"value", "volcano"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternSimpleMatch_wrongPattern_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("v[ou]l[", nil)
	if res := pattern.Match([]string{"value", "vol["}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_fromRootWithSlash(t *testing.T) {
	pattern := gitignore.ParsePattern("/value/vul?ano", nil)
	if res := pattern.Match([]string{"value", "vulkano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_withDomain(t *testing.T) {
	pattern := gitignore.ParsePattern("middle/tail/", []string{"value", "volcano"})
	if res := pattern.Match([]string{"value", "volcano", "middle", "tail"}, true); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_onlyMatchInDomain_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("volcano/tail", []string{"value", "volcano"})
	if res := pattern.Match([]string{"value", "volcano", "tail"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_fromRootWithoutSlash(t *testing.T) {
	pattern := gitignore.ParsePattern("value/vul?ano", nil)
	if res := pattern.Match([]string{"value", "vulkano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_fromRoot_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("value/vulkano", nil)
	if res := pattern.Match([]string{"value", "volcano"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_fromRoot_tooShort_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("value/vul?ano", nil)
	if res := pattern.Match([]string{"value"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_fromRoot_notAtRoot_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/value/volcano", nil)
	if res := pattern.Match([]string{"value", "value", "volcano"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_leadingAsterisks_atStart(t *testing.T) {
	pattern := gitignore.ParsePattern("**/*lue/vol?ano", nil)
	if res := pattern.Match([]string{"value", "volcano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_leadingAsterisks_notAtStart(t *testing.T) {
	pattern := gitignore.ParsePattern("**/*lue/vol?ano", nil)
	if res := pattern.Match([]string{"head", "value", "volcano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_leadingAsterisks_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("**/*lue/vol?ano", nil)
	if res := pattern.Match([]string{"head", "value", "Volcano", "tail"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_leadingAsterisks_isDir(t *testing.T) {
	pattern := gitignore.ParsePattern("**/*lue/vol?ano/", nil)
	if res := pattern.Match([]string{"head", "value", "volcano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_leadingAsterisks_isDirAtEnd(t *testing.T) {
	pattern := gitignore.ParsePattern("**/*lue/vol?ano/", nil)
	if res := pattern.Match([]string{"head", "value", "volcano"}, true); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_leadingAsterisks_isDir_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("**/*lue/vol?ano/", nil)
	if res := pattern.Match([]string{"head", "value", "Colcano"}, true); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_leadingAsterisks_isDirNoDirAtEnd_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("**/*lue/vol?ano/", nil)
	if res := pattern.Match([]string{"head", "value", "volcano"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_tailingAsterisks(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/vol?ano/**", nil)
	if res := pattern.Match([]string{"value", "volcano", "tail", "moretail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_tailingAsterisks_exactMatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/vol?ano/**", nil)
	if res := pattern.Match([]string{"value", "volcano"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_middleAsterisks_emptyMatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/**/vol?ano", nil)
	if res := pattern.Match([]string{"value", "volcano"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_middleAsterisks_oneMatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/**/vol?ano", nil)
	if res := pattern.Match([]string{"value", "middle", "volcano"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_middleAsterisks_multiMatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/**/vol?ano", nil)
	if res := pattern.Match([]string{"value", "middle1", "middle2", "volcano"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_middleAsterisks_isDir_trailing(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/**/vol?ano/", nil)
	if res := pattern.Match([]string{"value", "middle1", "middle2", "volcano"}, true); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_middleAsterisks_isDir_trailing_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/**/vol?ano/", nil)
	if res := pattern.Match([]string{"value", "middle1", "middle2", "volcano"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_middleAsterisks_isDir(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/**/vol?ano/", nil)
	if res := pattern.Match([]string{"value", "middle1", "middle2", "volcano", "tail"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_wrongDoubleAsterisk_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/*lue/**foo/vol?ano", nil)
	if res := pattern.Match([]string{"value", "foo", "volcano", "tail"}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_magicChars(t *testing.T) {
	pattern := gitignore.ParsePattern("**/head/v[ou]l[kc]ano", nil)
	if res := pattern.Match([]string{"value", "head", "volcano"}, false); res != gitignore.Exclude {
		t.Errorf("expected Exclude, found %v", res)
	}
}

func TestPatternGlobMatch_wrongPattern_noTraversal_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("**/head/v[ou]l[", nil)
	if res := pattern.Match([]string{"value", "head", "vol["}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}

func TestPatternGlobMatch_wrongPattern_onTraversal_mismatch(t *testing.T) {
	pattern := gitignore.ParsePattern("/value/**/v[ou]l[", nil)
	if res := pattern.Match([]string{"value", "head", "vol["}, false); res != gitignore.NoMatch {
		t.Errorf("expected NoMatch, found %v", res)
	}
}
