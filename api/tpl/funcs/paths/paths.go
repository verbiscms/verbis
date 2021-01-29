package paths

// Templates
//
// Returns the directory where page templates
// are stored.
//
// Example: {{ templatesPath }}
func (ns *Namespace) Templates() string {
	return ns.deps.Paths.Theme + ns.deps.Theme.TemplateDir
}

// Layouts
//
// Returns the directory where page layouts
// are stored.
//
// Example: {{ layoutsPath }}
func (ns *Namespace) Layouts() string {
	return ns.deps.Paths.Theme + ns.deps.Theme.LayoutDir
}
