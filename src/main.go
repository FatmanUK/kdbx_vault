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

var VERSION string = "badvalue"

// shut up, Gin
// no frillies
// no proxies?
func threadGinLoop(port string) {
	prefix := "/api/v1"
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.SetTrustedProxies(nil)
	addDatabaseRoutes(router, prefix)
	addToolsRoutes(router, prefix)
	addEntriesRoutes(router, prefix)
	addGroupsRoutes(router, prefix)
	router.Run("localhost:" + port)
}

func main() {
	dotv := DocOptVars{
		Name: APP_NAME,
		Version: VERSION,
		Port: KV_PORT,
	}
	ds := mustPrepareDocoptString(docoptString, dotv)
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
