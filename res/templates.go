package models

import (
	"cms/api/domain"
	"cms/api/helpers/paths"
	"cms/config"
	"encoding/json"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"os"
)

func discord() {
	master :=
		`
		Hello: {{ .names }}
		World: {{ .otherkey }}
		`

	data := make(map[string]string)
	data["names"] = "banana!"
	data["otherkey"] = "otherbanana!"

	masterTmpl, err := template.New("master").Parse(master)
	if err != nil {
		//log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, data); err != nil {
		//log.Fatal(err)
	}
}

package main

import (
"html/template"
"log"
"os"
)

type TemplateData struct {
	Fields map[string]string
}

func main() {

	templateString := `
Hello: {{ .Fields.names }}
World: {{ .Fields.other_key }}
`

	d := TemplateData{Fields: make(map[string]string)}
	d.Fields["names"] = "banana!"
	d.Fields["otherkey"] = "otherbanana!"

	masterTmpl, err := template.New("templateString").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, d); err != nil {
		log.Fatal(err)
	}
}

package templates

import (
"cms/api/domain"
"cms/api/helpers/paths"
"cms/config"
"encoding/json"
"github.com/foolin/goview"
"github.com/foolin/goview/supports/ginview"
"html/template"
)

type ViewData struct {
	Post		domain.Post
}


func Test() {
}

func GetViewData(post *domain.Post) *ViewData {

	fields := make(map[string]string)
	_ = json.Unmarshal(post.Fields, &fields)

	t := newFields(fields)


	c.server.HTMLRender = ginview.New(goview.Config{
		Root:      paths.Theme(),
		Extension: config.Template["file_extension"],
		Master:    "/layouts/main",
		Partials:  []string{},
		DisableCache: true,
		Funcs: template.FuncMap{
			"getField": func(field string) string {
				if _, found := fields[field]; found {
					return fields[field]
				} else {
					return ""
				}
			},
			"getFields": func() map[string]string {
				return fields
			},
			"hasField": func(field string) bool {
				if _, found := fields[field]; found {
					return true
				}
				return false
			},
			"test":
		},
	})

	return &ViewData{}
}

