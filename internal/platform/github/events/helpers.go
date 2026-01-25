package events

import (
	"fmt"
	"strings"
	"time"
)

var sgLoc = func() *time.Location {
	loc, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		return time.FixedZone("SGT", 8*60*60)
	}
	return loc
}()

// FormatUnixTimestamp formats a unix timestamp to a readable string.
func FormatUnixTimestamp(unixSec int64) string {
	if unixSec <= 0 {
		return ""
	}
	t := time.Unix(unixSec, 0).In(sgLoc)
	return t.Format("Mon, 2 Jan 2006, 3:04 PM") + " SGT"
}

// GetCurrentTimestamp returns the current timestamp as a formatted string.
func GetCurrentTimestamp() string {
	return time.Now().In(sgLoc).Format("Mon, 2 Jan 2006, 3:04 PM") + " SGT"
}

// FormatRFC3339Timestamp formats an RFC3339 timestamp to a readable string.
func FormatRFC3339Timestamp(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.In(sgLoc).Format("Mon, 2 Jan 2006, 3:04 PM") + " SGT"
}

// FormatRef extracts the branch/tag name from a full ref.
func FormatRef(ref string) string {
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

// ShortenSHA shortens a SHA to 7 characters.
func ShortenSHA(sha string) string {
	if len(sha) <= 7 {
		return sha
	}
	return sha[:7]
}

// EscapeText escapes a string for Telegram MarkdownV2 "normal text" context.
func EscapeText(s string) string {
	r := strings.NewReplacer(
		`\\`, `\\\\`,
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

// EscapeURL escapes the URL part inside [text](url) in MarkdownV2.
func EscapeURL(s string) string {
	r := strings.NewReplacer(
		`\`, `\\`,
		`)`, `\)`,
	)
	return r.Replace(s)
}

// FormatStringSlice formats a string slice for display.
func FormatStringSlice(slice []string) string {
	if len(slice) == 0 {
		return "(none)"
	}
	return strings.Join(slice, ", ")
}

// FormatInterfaceSlice formats an interface slice for display.
func FormatInterfaceSlice(slice *[]string) string {
	if slice == nil || len(*slice) == 0 {
		return "(none)"
	}
	result := make([]string, len(*slice))
	for i, v := range *slice {
		result[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(result, ", ")
}
