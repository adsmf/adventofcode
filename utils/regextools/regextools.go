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
