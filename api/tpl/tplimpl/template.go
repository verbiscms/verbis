package tplimpl

import (
	"github.com/ainsleyclark/verbis/api/deps"
)

// TemplateManager defines the service for executing and
// parsing Verbis templates. It's responsible for
// obtaining a template.FuncMap and Data to be
// used within the template.
type TemplateManager struct {
	deps *deps.Deps
}

// Config represents the options for passing
type Config struct {
	Root      string //view root
	Extension string //template extension
	Master    string //template master
}

// Creates a new TemplateManager
func New(d *deps.Deps) *TemplateManager {
	return &TemplateManager{}
}

func (c *Config) GetRoot() string {
	return c.Root
}

func (c *Config) GetExtension() string {
	return c.Extension
}

func (c *Config) GetMaster() string {
	return c.Master
}