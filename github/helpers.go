package github

import (
	"strings"
	"time"
)


// unixSec is seconds since epoch (GitHub repository.pushed_at).
// Example output: "Thu, 18 Dec 2025, 1:03 AM SGT"
func formatUnixTimestamp(unixSec int64) string {
	if unixSec <= 0 {
		return ""
	}

	loc, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		loc = time.FixedZone("SGT", 8*60*60)
	}

	t := time.Unix(unixSec, 0).In(loc)
	return t.Format("2 Jan 2006, Mon, 3:04 PM") + " SGT"
}


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

// escapeText escapes a string for Telegram MarkdownV2 "normal text" context.
func escapeText(s string) string {
	// Escape backslash first, then the rest.
	r := strings.NewReplacer(
		`\\`, `\\\\`, // if your input can already contain escapes, keep this; otherwise use "\" -> "\\"
		`\`, `\\`,
		`_`, `\_`,
		`*`, `\*`,
		`[`, `\[`,
		`]`, `\]`,
		`(`, `\(`,
		`)`, `\)`,
		`~`, `\~`,
		"`", "\\`",
		`>`, `\>`,
		`#`, `\#`,
		`+`, `\+`,
		`-`, `\-`,
		`=`, `\=`,
		`|`, `\|`,
		`{`, `\{`,
		`}`, `\}`,
		`.`, `\.`,
		`!`, `\!`,
	)
	return r.Replace(s)
}


// escapeURL escapes the URL part inside [text](url) in MarkdownV2.
func escapeURL(s string) string {
	r := strings.NewReplacer(
		`\`, `\\`,
		`)`, `\)`,
	)
	return r.Replace(s)
}