package rex

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerators(t *testing.T) {
	for _, test := range []struct {
		expression *regexp.Regexp
		expected   string
	}{
		{
			Match("something"),
			"something",
		},
		{
			AnyLit("amb", "blu", "brn"),
			"amb|blu|brn",
		},
		{
			Anchor(
				Literal("prefix."),
				Any(Match("1?2"), Match("34*"), Match("2137")),
			),
			`^(?:prefix\.(?:1?2|34*|2137))$`,
		},
		{
			NumberBetween(0, 9),
			"(?:[0-9])",
		},
		{
			NumberBetween(8, 12),
			"(?:[8-9]|1[0-2])",
		},
		{
			NumberBetween(12, 17),
			"(?:1[2-7])",
		},
		{
			NumberBetween(1920, 2002),
			"(?:19[2-9][0-9]|200[0-2])",
		},
		{
			NumberBetween(2010, 2020),
			"(?:201[0-9]|2020)",
		},
		{
			NumberBetween(150, 193),
			"(?:1[5-8][0-9]|19[0-3])",
		},
		{
			NumberBetween(7, 218),
			"(?:[7-9]|[1-9][0-9]|1[0-9][0-9]|20[0-9]|21[0-8])",
		},
	} {

		assert.Equal(t, test.expected, test.expression.String())
	}
}
