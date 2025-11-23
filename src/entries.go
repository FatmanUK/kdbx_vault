package main

import (
//	"os"
//	"net/http"
	"github.com/gin-gonic/gin"
//	kdbx "github.com/tobischo/gokeepasslib/v3"
)

func addEntriesRoutes(r *gin.Engine, pre string) {
	r.GET(pre + "/entries", getEntries)
	r.PUT(pre + "/entries", putEntries)
	r.GET(pre + "/entries/:id", getEntryById)
	r.POST(pre+ "/entries/:id", postEntryById)
	r.DELETE(pre + "/entries/:id", deleteEntryById)
}

/*
curl http://localhost:8192/api/v1/entries -H "X-Kdbx-B64Password: YWJjZDEyMzQjCg==" --no-progress-meter | jq '.Values[] | select (.Key=="Password")'
*/

// -H "X-Kdbx-B64Password: YWJjZDEyMzQjCg=="
func getEntries(c *gin.Context) { // list entries
	logs <- "Get Entries"
/*
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	db = kdbx.NewDatabase()
	db.Credentials = kdbx.NewPasswordCredentials(pw)
	_ = kdbx.NewDecoder(file).Decode(db)
	db.UnlockProtectedEntries()
	defer db.LockProtectedEntries()
	entry := db.Content.Root.Groups[0].Groups[0].Entries[0]

	c.IndentedJSON(http.StatusOK, entry)
*/
}

/*

    // Note: This is a simplified example and the groups and entries will depend on the specific file.
    // bound checking for the slices is recommended to avoid panics.
    entry := db.Content.Root.Groups[0].Groups[0].Entries[0]
    fmt.Println(entry.GetTitle())
    fmt.Println(entry.GetPassword())
*/

//func postEntries(c *gin.Context) {
//}

func putEntries(c *gin.Context) { // add new entry
	logs <- "Put Entries"
}

func getEntryById(c *gin.Context) { // return entry
	logs <- "Get Entry"
/*router.GET("/user/:id", func(c *gin.Context) {
    // a GET request to /user/john
    id := c.Param("id") // id == "john"
    // a GET request to /user/john/
    id := c.Param("id") // id == "/john/"
})*/
//	id := c.Param("id")
}

func postEntryById(c *gin.Context) { // update entry
	logs <- "Post Entry"
//	id := c.Param("id")
}

func deleteEntryById(c *gin.Context) { // delete entry
	logs <- "Delete Entry"
//	id := c.Param("id")
}
