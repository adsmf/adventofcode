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
	} {
		assert.Equal(t, test.expected, test.expression.String())
	}
}
