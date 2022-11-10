package util

import (
	"github.com/google/uuid"
	"os"
	"text/template"
)

func TemplateParseFile(templateFilePath, targetFilePath string, data interface{}) (err error) {
	tpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return
	}
	targetFile, err := os.Open(targetFilePath)
	if err != nil {
		return
	}
	defer targetFile.Close()

	return tpl.Execute(targetFile, data)
}

func TemplateParseString(templateString, targetFilePath string, data interface{}) (err error) {
	tpl, err := template.New(uuid.NewString()).Parse(templateString)
	if err != nil {
		return
	}
	targetFile, err := os.Open(targetFilePath)
	if err != nil {
		return
	}
	defer targetFile.Close()

	return tpl.Execute(targetFile, data)
}
