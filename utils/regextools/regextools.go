package rex

import (
	"fmt"
	"regexp"
	"strings"
)

var Match = regexp.MustCompile

// Basics
func Literal(s string) *regexp.Regexp {
	return Match(regexp.QuoteMeta(s))
}
func Join(expressions ...*regexp.Regexp) *regexp.Regexp {
	combined := ""
	for _, exp := range expressions {
		combined += exp.String()
	}
	return Match(combined)
}

// Groups
func Group(expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match("(?:" + Join(expressions...).String() + ")")
}
func Anchor(expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match("^" + Group(expressions...).String() + "$")
}
func Any(expressions ...*regexp.Regexp) *regexp.Regexp {
	opts := []string{}
	for _, exp := range expressions {
		opts = append(opts, exp.String())
	}
	return Join(Group(Match(strings.Join(opts, "|"))))
}
func AnyLit(options ...string) *regexp.Regexp {
	return Join(Match(strings.Join(options, "|")))
}

func NumberBetween(min, max int) *regexp.Regexp {
	groups := regexForRange(min, max+1)
	return Group(Match(strings.Join(groups, "|")))
}
func regexForRange(min, max int) []string {
	if max <= min {
		return []string{}
	}

	prefix := fmt.Sprintf("%d", min)
	suffix := fmt.Sprintf("%d", max)
	minUnits := min % 10
	if int(min/10) == int((max-1)/10) {
		maxUnits := (max - 1) % 10
		return []string{prefix[:len(prefix)-1] + digitRangeGroup(minUnits, maxUnits)}
	}

	prefixList := []string{}
	if minUnits != 0 {
		prefixList = append(prefixList, prefix[:len(prefix)-1]+digitRangeGroup(minUnits, 9))
	}

	suffixList := []string{}
	maxUnits := max % 10
	if maxUnits != 0 {
		suffixList = append(suffixList, suffix[:len(suffix)-1]+digitRangeGroup(0, maxUnits-1))
	}

	midList := []string{}
	for _, mid := range regexForRange(int((min+9)/10), int(max/10)) {
		midList = append(midList, mid+"[0-9]")
	}
	result := append(prefixList, midList...)
	result = append(result, suffixList...)
	return result
}
func digitRangeGroup(min, max int) string {
	if min == max {
		return fmt.Sprintf("%d", min)
	}
	return fmt.Sprintf("[%d-%d]", min, max)
}

// Repetition modifiers
func Optional(expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(expressions...).String() + "?")
}
func ZeroPlus(expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(expressions...).String() + "*")
}
func Repeated(expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(expressions...).String() + "+")
}
func Times(reps int, expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(expressions...).String() + fmt.Sprintf("{%d}", reps))
}
func TimesAtLeast(reps int, expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(expressions...).String() + fmt.Sprintf("{%d,}", reps))
}
func TimesBetween(minReps, maxReps int, expressions ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(expressions...).String() + fmt.Sprintf("{%d,%d}", minReps, maxReps))
}

// Character groups
var Digit = Match("[0-9]")
var HexChar = Match("[0-9a-f]")
