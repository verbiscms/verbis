package minify

const (
	// HTML mime type for comparison
	HTML = "text/html"
	// CSS mime type for comparison
	CSS = "text/css"
	// Javascript mime type for comparison
	Javascript = "application/javascript"
	// SVG mime type for comparison
	SVG = "image/svg+xml"
	// JSON mime type for comparison
	JSON = "application/json"
	// XML mime type for comparison
	XML = "text/xml"
)

var (
	// The default configuration for when none is passed
	defaultConfig = Config{
		MinifyHTML: true,
		MinifyCSS:  true,
		MinifyJS:   true,
		MinifySVG:  true,
		MinifyJSON: true,
		MinifyXML:  true,
	}
	// htmlMime is the default for HTML regexp for the pkg minifier
	htmlMime = `text/html`
	// cssMime is the default for CSS regexp for the pkg minifier
	cssMime = `text/css`
	// jsMime is the default for JS regexp for the pkg minifier
	jsMime = `^(application|text)/(x-)?(java|ecma)script$`
	// svgMime is the default for SVG regexp for the pkg minifier
	svgMime = `image/svg+xml`
	// jsonMime is the default for JSON regexp for the pkg minifier
	jsonMime = `[/+]json$`
	// xmlMime is the default for XML regexp for the pkg minifier
	xmlMime = `[/+]xml$`
)

// Config represents the options for minifying output.
type Config struct {
	MinifyHTML bool
	MinifyCSS  bool
	MinifyJS   bool
	MinifySVG  bool
	MinifyJSON bool
	MinifyXML  bool
}
