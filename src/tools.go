package main

import (
	"github.com/gin-gonic/gin"
	//kdbx "github.com/tobischo/gokeepasslib/v3"
)

func addToolsRoutes(r *gin.Engine, pre string) {
	r.GET(pre + "/tools", getTools)
	r.GET(pre + "/tools/password", getToolsPassword)
	r.GET(pre + "/tools/password/list", getToolsPasswordList)
}

func getTools(c *gin.Context) { // list of tools
	logs <- "Get Tools"
}

func getToolsPassword(c *gin.Context) { // new password
	logs <- "Get Password"
}

func getToolsPasswordList(c *gin.Context) { // batch of new passwords
	logs <- "Get Batch Passwords"
}
