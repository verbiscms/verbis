package attributes

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// Body
//
// Returns class names for the body element. Includes the
// resource, page ID, page title, page template, page
// layout and if the user is logged in or not.
//
// Example: `{{ body }}`
// Returns: `page page-id-4 page-title page-template-news-archive page-layout-main logged-in` (for example)
func (ns *Namespace) Body() string {
	body := new(bytes.Buffer)
	p := ns.post.Post

	// Resource, writes page if no resource (e.g. page)
	if p.Resource == nil {
		body.WriteString("page ")
	} else {
		body.WriteString(fmt.Sprintf("%s ", *p.Resource))
	}

	// Page ID (e.g. page-id-2)
	body.WriteString(fmt.Sprintf("page-id-%v ", p.Id))

	// Title
	body.WriteString(fmt.Sprintf("page-title-%s ", cssValidString(p.Title)))

	// Page template (e.g. page-template-test)
	body.WriteString(fmt.Sprintf("page-template-%s ", cssValidString(p.PageTemplate)))

	// Page Layout (e.g page-layout-test)
	body.WriteString(fmt.Sprintf("page-layout-%s", cssValidString(p.PageLayout)))

	// Logged in (e.g. logged-in) if auth
	//if ns.auth.Auth() {
	//	body.WriteString(" logged-in")
	//}

	return body.String()
}

// lang
//
// Returns language attributes set in the options for
// use with the `<html lang="">` attribute.
//
// Example: `{{ lang }}`
// Returns: 'en-gb` (for example)
func (ns *Namespace) Lang() string {
	return ns.deps.Options.GeneralLocale
}

// cssValidString
//
// Strips all special characters from the given string
// and replaces forward slashes with hyphens & spaces
// with dashes for a valid CSS class.
//
// Example:
// `My Page !Template` would return `my-page-template`.
func cssValidString(str string) string {
	r := regexp.MustCompile("[^A-Za-z0-9\\s-/]")

	str = r.ReplaceAllString(str, "")
	str = strings.ReplaceAll(str, "/", "-")
	str = strings.ReplaceAll(str, " ", "-")

	return strings.ToLower(str)
}
