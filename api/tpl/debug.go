package tpl

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"html/template"
	"reflect"
	"sort"
	"time"
)

func  (t *TemplateManager) debug(i interface{}) string {
	return fmt.Sprintf("%+v\n", i)
}

func (t *TemplateManager) dd(i interface{}) (template.HTML, error) {
	d := debug{
		Bytes: new(bytes.Buffer),
	}
	str, err := d.Dump(i)
	return template.HTML(str), err
}

type tester struct {
	test string
	testStruct struct{
		testinner string
	}
}

type debugger interface {
	Dump(i interface{}) (string, error)
}

// Have a struct!

type debug struct {
	Bytes *bytes.Buffer
}

func (d *debug) Dump(i interface{}) (string, error) {

	d.Bytes.WriteString(`
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


	d.test("", "", -1, i)
	d.Bytes.WriteString("</pre>")

	file, err := template.New("debug").Parse(d.Bytes.String())
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = file.Execute(&tpl, nil)
	if err != nil {
		return "", fmt.Errorf("Unable to execute partial file: %v", err)
	}

	return tpl.String(), nil
}

func (d *debug) Sort(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0

	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	return keys
}

// Indent
//
//
func (d *debug) Indent(lvl int) *debug {
	for i := 0; i < lvl; i++ {
		d.Bytes.WriteString("    ")
	}
	return d
}

// LineBreak
//
//
func (d *debug) LineBreak() *debug {
	d.Bytes.WriteString("\n")
	return d
}

// Write
//
//
func (d *debug) Write(key string, typ string, val string) *debug {
	d.Bytes.WriteString("+ ")
	d.Bytes.WriteString(fmt.Sprintf("<span class=\"sf-dump-name\">%s</span> ", key))
	d.Bytes.WriteString(fmt.Sprintf("<span class=\"sf-dump-type\">%s</span> ", typ))
	d.Bytes.WriteString(fmt.Sprintf("<span class=\"sf-dump-value\">\"%s\"</span> ", val))
	return d
}

// test
//
//
func (d *debug) test(key string, typ string, lvl int, e interface{}) {

	lvl = lvl + 1

	val := reflect.ValueOf(e)
	if val.Kind() == reflect.Ptr {
		e = reflect.Indirect(val)
	}

	if e == nil {
		d.Indent(lvl).Write(key, typ, "nil")
		return
	}

	if reflect.TypeOf(e).Kind() == reflect.Struct {
		d.DeepFields(e)
	//	d.Struct(typ, lvl, e)
	}

	switch v := e.(type) {
	case int, int8, int16, int32, int64:
		d.Indent(lvl).Write(key, typ, fmt.Sprintf("%d", v))
	case string:
		d.Indent(lvl).Write(key, typ, fmt.Sprintf("%s", v))
	case bool:
		d.Indent(lvl).Write(key, typ, fmt.Sprintf("%v", v))
	case time.Time:
		d.Indent(lvl).Write(key, typ, fmt.Sprintf("%v", v))
	case map[string]interface{}:
		d.Indent(lvl).Bytes.WriteString(fmt.Sprintf("%s {\n", key))
		for _, k := range d.Sort(v) {
			d.test(k, reflect.TypeOf(v[k]).String(), lvl, v[k])
			d.LineBreak()
		}
		d.Indent(lvl).Bytes.WriteString("}")
	}
}


// Struct
//
//
func (d *debug) Struct(typ string, lvl int, e interface{}) {
	numField := reflect.TypeOf(e).NumField()

	d.Indent(lvl).Bytes.WriteString(fmt.Sprintf("%s {\n", typ))
	d.Bytes.WriteString(fmt.Sprintf("<samp data-depth=\"%d\">", numField))

	for i := 0 ; i < numField ; i++ {
		field := reflect.TypeOf(e).Field(i)
		val := reflect.ValueOf(e).Field(i)

		//if val.Kind() == reflect.Struct {
		//	continue
		//}


		if val.Kind() == reflect.Ptr {
			val = reflect.Indirect(val)
		}

		if val.Kind() == reflect.Map {
			color.Yellow.Println(val)
			color.Yellow.Println(field.Name)
		}

		if !val.IsValid() {
			d.Indent(lvl + 1).Bytes.WriteString(fmt.Sprintf(`<span class="sf-dump-name">%s</span> `, field.Name))
			str := fmt.Sprintf(`<span class="sf-dump-type">%s</span> <span class="sf-dump-value">nil</span>`, field.Type)
			d.Bytes.WriteString(str)
			d.LineBreak()
			continue
		}

		if val.CanInterface() {
			d.Indent(lvl + 1).Bytes.WriteString(fmt.Sprintf(`<span class="sf-dump-name">%s</span> `, field.Name))
			str := fmt.Sprintf(`<span class="sf-dump-type">%T</span> <span class="sf-dump-value">%[1]v</span>`, val.Interface())
			d.Bytes.WriteString(str)
			d.LineBreak()
		}
	}

	d.Indent(lvl).Bytes.WriteString(fmt.Sprintf("</samp>"))
	d.Bytes.WriteString(fmt.Sprintf("}"))
}

func (d *debug) DeepFields(iface interface{})  {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			d.DeepFields(v.Interface())
		default:
			d.Indent(1).Bytes.WriteString(fmt.Sprintf(`<span class="sf-dump-name">%s</span> `, ift.Field(i).Name))
			str := fmt.Sprintf(`<span class="sf-dump-type">%T</span> <span class="sf-dump-value">%[1]v</span>`, v.Interface())
			d.Bytes.WriteString(str)
			d.LineBreak()
		}
	}
}