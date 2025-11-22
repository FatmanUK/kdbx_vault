package main

import (
	"github.com/gin-gonic/gin"
	//kdbx "github.com/tobischo/gokeepasslib/v3"
)

type Tool struct {
	Desc Descriptor
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
