// Package strutil provides some string utilities.
package strutil

// FirstNonEmpty returns the first non-empty string.
func FirstNonEmpty(strings ...string) string {
	for _, s := range strings {
		if s != "" {
			return s
		}
	}

	return ""
}
