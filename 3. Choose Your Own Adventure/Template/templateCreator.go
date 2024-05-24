package template

import (
	"html/template"
	"os"
)

func TemplateCreate(data any) {
	var tmplFile = "story.tmpl"
	os.Chdir("Template")
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
