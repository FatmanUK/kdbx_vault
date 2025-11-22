package main

import (
	"os"
	"bytes"
	"strconv"
	"text/template"
	"github.com/gin-gonic/gin"
	kdbx "github.com/tobischo/gokeepasslib/v3"
)

// These are set during build.
var BUILD_MODE string
var APP_NAME string
var VERSION string
var KV_PORT string

type Descriptor struct {
	Ident string `json:"id"`
	Type string `json:"type"`
	Metadata map[string]interface{} `json:"metadata"`
}

type Config struct {
}

type DocOptTemplateVar struct {
	Name string
	Version string
	Port string
}

var logs chan string
var filename string

func intFromString(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

//var file *os.File
var db *kdbx.Database

/*
Use cases:
KDBX_PASSWORD='abcdfegh1234#'
KDBX_PORT: '8192'
curl http://localhost:${KDBX_PORT}/api/v1/...
*/

func getSingleHeader(c *gin.Context, headerName string) string {
	header := c.Request.Header[headerName]
	l := len(header)
	if l == 0 {
		panic("Too few '" + headerName + "' headers provided")
		return ""
	}
	if l > 1 {
		panic("Too many '" + headerName + "' headers provided")
		return ""
	}
	return header[0]
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func prepareDocoptString(ds string, dotv DocOptTemplateVar) string {
	docoptTemplate := template.New("docoptTemplate")
	docoptTemplate = template.Must(docoptTemplate.Parse(ds))
	var wr bytes.Buffer
	err := docoptTemplate.Execute(&wr, dotv)
	if err != nil {
		panic(err)
	}
	return string(wr.Bytes())
}

func mustOpenFile(fn string) *os.File {
	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	return file
}

func openKdbx(file *os.File, pw string) *kdbx.Database {
	db := kdbx.NewDatabase()
	db.Credentials = kdbx.NewPasswordCredentials(pw)
	_ = kdbx.NewDecoder(file).Decode(db)
	return db
}
