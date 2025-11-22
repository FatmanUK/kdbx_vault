package main

import (
	"log"
	"strings"
	"github.com/gin-gonic/gin"
	doh "github.com/FatmanUK/fatgo/docopt_helpers"
)

const docoptString = `{{ .Name }} v{{ .Version }}

Usage:
  {{ .Name }}
  {{ .Name }} -f <kdbx> [-p <port>]
  {{ .Name }} -h | --help
  {{ .Name }} --version

Options:
  -h --help  Show this screen
  --version  Show version
  -f <kdbx>  Load KDBX database
  -p <port>  Localhost port [default: {{ .Port }}]
`

func threadGinLoop(port string) {
	prefix := "/api/v1"
	gin.SetMode(gin.ReleaseMode) // shut up, Gin
	router := gin.New() // no frillies
	router.SetTrustedProxies(nil) // no proxies?

	router.GET(prefix + "/database", getDatabase)
	router.POST(prefix + "/database", postDatabase)
/*
	router.GET(prefix + "/tools", getTools)
	router.GET(prefix + "/tools/password", getToolsPassword)
	router.GET(prefix + "/tools/password/list", getToolsPasswordList)

	router.GET(prefix + "/entries", getEntries)
	router.PUT(prefix + "/entries", putEntries)
	router.GET(prefix + "/entries/:id", getEntryById)
	router.POST(prefix + "/entries/:id", postEntryById)
	router.DELETE(prefix + "/entries/:id", deleteEntryById)

	router.GET(prefix + "/groups", getGroups)
	router.PUT(prefix + "/groups", putGroups)
	router.GET(prefix + "/groups/:id", getGroupById)
	router.POST(prefix + "/groups/:id", postGroupById)
	router.DELETE(prefix + "/groups/:id", deleteGroupById)
*/
	router.Run("localhost:" + port)
}

func main() {
	dotv := DocOptTemplateVar{
		Name: APP_NAME,
		Version: VERSION,
		Port: KV_PORT,
	}
	ds := prepareDocoptString(docoptString, dotv)
	appVer := dotv.Name + " v" + dotv.Version
	args, err := doh.NoExitParser.ParseArgs(ds, nil, appVer)
	if err != nil {
		panic(err)
	}
	filename = args["-f"].(string)
	logs = make(chan string)
	go threadGinLoop(args["-p"].(string))
	for msg := range logs {
		msgs := strings.Split(msg, "\\n")
		for _, m := range removeEmptyStrings(msgs) {
			log.Println(m)
		}
	}
}
