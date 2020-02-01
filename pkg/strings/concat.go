// +build go1.10

package strings

import gostrings "strings"

// Concat concatenates all passed strings.
//
// Deprecated: use go.octolab.org/strings.Concat instead.
func Concat(strings ...string) string {
	b := gostrings.Builder{}
	for _, str := range strings {
		b.WriteString(str)
	}
	return b.String()
}
