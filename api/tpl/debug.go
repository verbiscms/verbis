package tpl

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/tpl/funcs/debug"
)

func (t *TemplateManager) debug(i interface{}) string {
	return fmt.Sprintf("%+v\n", i)
}

func (t *TemplateManager) dd(i interface{}) (string, error) {
	return debug.Dump(i), nil
}
