package main

import (
	"github.com/gin-gonic/gin"
	//kdbx "github.com/tobischo/gokeepasslib/v3"
)

type Group struct {
	Desc Descriptor
}

func getGroups(c *gin.Context) { // list groups
	logs <- "Get Groups"
}

func putGroups(c *gin.Context) { // add group
	logs <- "Put Groups"
}

func getGroupById(c *gin.Context) { // return group data
	logs <- "Get Group"
}

func postGroupById(c *gin.Context) { // update group
	logs <- "Post Group"
}

func deleteGroupById(c *gin.Context) { // delete group
	logs <- "Delete Group"
}
