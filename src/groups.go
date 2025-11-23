package main

import (
	"github.com/gin-gonic/gin"
	//kdbx "github.com/tobischo/gokeepasslib/v3"
)

func addGroupsRoutes(r *gin.Engine, pre string) {
	r.GET(pre + "/groups", getGroups)
	r.PUT(pre + "/groups", putGroups)
	r.GET(pre + "/groups/:id", getGroupById)
	r.POST(pre+ "/groups/:id", postGroupById)
	r.DELETE(pre + "/groups/:id", deleteGroupById)
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
