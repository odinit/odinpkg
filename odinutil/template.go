package odinutil

import (
	"github.com/google/uuid"
	"io"
	"text/template"
)

func TemplateParseFile(tp string, wr io.Writer, data interface{}) (err error) {
	tpl, err := template.ParseFiles(tp)
	if err != nil {
		return
	}

	return tpl.Execute(wr, data)
}

func TemplateParseString(ts string, wr io.Writer, data interface{}) (err error) {
	tpl, err := template.New(uuid.NewString()).Parse(ts)
	if err != nil {
		return
	}

	return tpl.Execute(wr, data)
}
