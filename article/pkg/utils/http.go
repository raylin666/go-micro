package utils

import "strings"

const (
	httpBaseContentType = "application"
)

// HttpContentType returns the content-type with base prefix.
func HttpContentType(subtype string) string {
	return strings.Join([]string{httpBaseContentType, subtype}, "/")
}

// HttpContentSubtype returns the content-subtype for the given content-type.  The
// given content-type must be a valid content-type that starts with
// but no content-subtype will be returned.
// according rfc7231.
// contentType is assumed to be lowercase already.
func HttpContentSubtype(contentType string) string {
	left := strings.Index(contentType, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentType, ";")
	if right == -1 {
		right = len(contentType)
	}
	if right < left {
		return ""
	}
	return contentType[left+1 : right]
}
