package tpl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetRoot(t *testing.T) {
	c := Config{Root: "test"}
	got := c.GetRoot()
	assert.Equal(t, "test", got)
}

func TestConfig_GetExtension(t *testing.T) {
	c := Config{Extension: "test"}
	got := c.GetExtension()
	assert.Equal(t, "test", got)
}

func TestConfig_GetMaster(t *testing.T) {
	c := Config{Master: "test"}
	got := c.GetMaster()
	assert.Equal(t, "test", got)
}