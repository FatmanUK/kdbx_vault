package main

import (
	"bytes"
	"text/template"
)

type DocOptVars struct {
	Name string
	Version string
	Port string
}

func mustPrepareDocoptString(ds string, dotv DocOptVars) string {
	docoptTemplate := template.New("docoptTemplate")
	docoptTemplate = template.Must(docoptTemplate.Parse(ds))
	var wr bytes.Buffer
	err := docoptTemplate.Execute(&wr, dotv)
	if err != nil {
		panic(err)
	}
	return string(wr.Bytes())
}
