package tpl

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

func  (t *TemplateManager) debug(i interface{}) string {
	return fmt.Sprintf("%+v\n", i)
}

// Have a struct!

type dump struct {
	bytes *bytes.Buffer

}

// or
type dd *bytes.Buffer

// or both


func (t *TemplateManager) dd(i interface{}) (template.HTML, error) {


	var b bytes.Buffer
	b.WriteString(`
<style>
	.pre.sf-dump {
		background-color: #18171B;
		color: #FF8400;
		line-height: 1.4em;
		font: 12px Menlo, Monaco, Consolas, monospace;
		word-wrap: break-word;
		white-space: pre-wrap;
		position: relative;
		z-index: 99999;
		word-break: break-all;
    	white-space: pre-wrap;
		padding: 5px;
		overflow: initial !important;

	}
	.pre.sf-dump .sf-dump-name {
     	color: #fff;
	}
	.pre.sf-dump .sf-dump-public {
     	color: #fff;
	}
	.pre.sf-dump .sf-dump-value {
		color: #56DB3A;
		font-weight: bold;
	}
</style>
<pre class="pre sf-dump">`)


	test(&b, "", "", -1, i)
	b.WriteString("</pre>")

	file, err := template.New("debug").Parse(b.String())
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = file.Execute(&tpl, nil)
	if err != nil {
		return "", fmt.Errorf("Unable to execute partial file: %v", err)
	}

	return template.HTML(tpl.String()), nil
}


func test(b *bytes.Buffer, key string, typ string, lvl int, e interface{}) {
	lvl = lvl + 1

	if e == nil {
		b.WriteString("<i>Nil</i>")
	}

	if reflect.TypeOf(e).Kind() == reflect.Struct {
		handleStruct(b, e)
	}

	switch e.(type) {
	case int, int8, int16, int32, int64:
		level(b, lvl)
		write(b, key, typ, fmt.Sprintf("%d", e.(int)))
	case string:
		level(b, lvl)
		write(b, key, typ, fmt.Sprintf("%s", e.(string)))
	case bool:
		level(b, lvl)
		write(b, key, typ, fmt.Sprintf("%v", e.(bool)))
	case map[string]interface{}:
		m := e.(map[string]interface{})
		level(b, lvl)
		b.WriteString(fmt.Sprintf("{ %s\n", key))
		for k, t := range m {
			test(b, k, reflect.TypeOf(t).String(), lvl, t)
			b.WriteString("\n")
		}
		level(b, lvl)
		b.WriteString(fmt.Sprintf("}"))
	}
	return
}

func level(b *bytes.Buffer, lvl int) {
	for i := 0; i < lvl; i++ {
		b.WriteString("    ")
	}
}

func write(b *bytes.Buffer, key string, typ string, val string) {
	b.WriteString("+ ")
	b.WriteString(fmt.Sprintf("<span class=\"sf-dump-name\">%s</span> ", key))
	b.WriteString(fmt.Sprintf("<span class=\"sf-dump-type\">%s</span> ", typ))
	b.WriteString(fmt.Sprintf("<span class=\"sf-dump-value\">%s</span> ", val))
}


func handleStruct(b *bytes.Buffer, e interface{}) {
	v := reflect.TypeOf(e)
	typ := reflect.ValueOf(e).Type()
	numField := reflect.TypeOf(e).NumField()

	b.WriteString(fmt.Sprintf("%s {\n", typ))
	b.WriteString(fmt.Sprintf("<samp data-depth=\"%d\">", numField))

	for i := 0 ; i < numField ; i++ {
		field := v.Field(i)
		value := reflect.ValueOf(e).Field(i)

		b.WriteString("    + ")
		b.WriteString(fmt.Sprintf("<span class=\"sf-dump-name\">%s</span> ", field.Name))
		b.WriteString(fmt.Sprintf("<span class=\"sf-dump-type\">%s</span> ", field.Type))

		if value.Kind() == reflect.Ptr {
			fmt.Println(value.Elem())
			if value.IsNil() {
				b.WriteString("<span class=\"sf-dump-value\">nil</span>\n")
			}
			continue
		}

		b.WriteString(fmt.Sprintf("<span class=\"sf-dump-value\">\"%s\"</span>\n", value.String()))
	}

	b.WriteString(fmt.Sprintf("</samp>"))
	b.WriteString(fmt.Sprintf("}"))
}