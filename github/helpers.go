package github

import (
	"strings"
)

func formatRef(ref string) string {
	if ref == "" {
		return ""
	}
	if strings.HasPrefix(ref, "refs/heads/") {
		return strings.TrimPrefix(ref, "refs/heads/")
	}
	if strings.HasPrefix(ref, "refs/tags/") {
		return strings.TrimPrefix(ref, "refs/tags/")
	}
	if strings.HasPrefix(ref, "refs/") {
		return strings.TrimPrefix(ref, "refs/")
	}
	return ref
}

func shortenSHA(sha string) string {
	if len(sha) <= 7 {
		return sha
	}
	return sha[:7]
}