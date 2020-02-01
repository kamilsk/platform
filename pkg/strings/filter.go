package strings

// FirstValid returns a first non-empty string.
//
// Deprecated: use go.octolab.org/strings.FirstNotEmpty instead.
func FirstValid(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}

// NotEmpty filters empty strings in-place.
//
// Deprecated: use go.octolab.org/strings.NotEmpty instead.
func NotEmpty(strings []string) []string {
	filtered := strings[:0]
	for _, str := range strings {
		if str != "" {
			filtered = append(filtered, str)
		}
	}
	return filtered
}
