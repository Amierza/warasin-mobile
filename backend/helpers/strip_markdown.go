package helpers

import "regexp"

func StripMarkdown(md string) string {
	// Hapus tanda heading #
	re := regexp.MustCompile(`(?m)^#+\s*`)
	md = re.ReplaceAllString(md, "")

	// Hapus backtick pada code block & inline code
	md = regexp.MustCompile("(?s)```.*?```").ReplaceAllString(md, "")
	md = regexp.MustCompile("`([^`]*)`").ReplaceAllString(md, "$1")

	// Ubah [text](link) jadi text (link)
	md = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`).ReplaceAllString(md, "$1 ($2)")

	// Ubah \n\n jadi 1 newline
	md = regexp.MustCompile(`\n{2,}`).ReplaceAllString(md, "\n")

	return md
}
