// Package embed provides functionality for matching and replacing URLs based on
// a specific platform.
//
// It allows the creation of an Embed instance that can check if a given URL
// matches the pattern for a specific platform's website.

package embed

import (
	"fmt"
	"regexp"
)

// Embed represents the structure for URL pattern matching and replacement.
type Embed struct {
	pattern  *regexp.Regexp
	replacer string
}

// New creates a new Embed instance with a specified platform and replacer string.
func New(platform string, replacer string) *Embed {
	return &Embed{
		pattern:  regexp.MustCompile(fmt.Sprintf(`(?i)^(https?://)?((www|vm)\.)?%s\.com`, platform)),
		replacer: replacer,
	}
}

// Check verifies whether the provided URL matches the predefined pattern.
func (e Embed) Check(url string) bool {
	return e.pattern.MatchString(url)
}

// Replace substitutes the matched part of the URL with the replacer string.
func (e Embed) Replace(url string) string {
	return e.pattern.ReplaceAllString(url, fmt.Sprintf("${1}${2}%s", e.replacer))
}
