package templates

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func TestPartial_Wrong_Path(t *testing.T) {
	tf := TemplateFunctions{}
	tpl := `{{ partial "wrongval" }}`
	temp, _ := template.New("").Funcs(tf.GetFunctions()).Parse(tpl)

	var bytes bytes.Buffer
	err := temp.Execute(&bytes, nil)

	assert.Error(t, err, fmt.Errorf("No file exists with the path: wrongval"))
	var err error
}

//func TestPartial_Didnt_Execute(t *testing.T) {
//
//	tf := TemplateFunctions{}
//	tpl := `{{ partial "wrongval" }}`
//	temp, _ := template.New("").Funcs(tf.GetFunctions()).Parse(tpl)
//}