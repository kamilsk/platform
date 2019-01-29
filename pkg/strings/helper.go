package strings

// FirstValid returns a first non-empty string.
func FirstValid(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}
