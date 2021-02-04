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
	htmlMime = `text/html`
	cssMime = `text/css`
	jsMime = `^(application|text)/(x-)?(java|ecma)script$`
	svgMime = `image/svg+xml`
	jsonMime = `[/+]json$`
	xmlMime =`[/+]xml$`
)

// Config represents the options for minifying output.
type Config struct {
	MinifyHTML bool
	MinifyCSS bool
	MinifyJS bool
	MinifySVG bool
	MinifyJSON bool
	MinifyXML bool
}