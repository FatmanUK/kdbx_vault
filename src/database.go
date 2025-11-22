package main

import (
	//"os"
	"github.com/gin-gonic/gin"
	//kdbx "github.com/tobischo/gokeepasslib/v3"
)

type Database struct {
	Desc Descriptor
}

func getDatabase(c *gin.Context) { // return database metadata
	logs <- "Get Database"

	b64_pw := getSingleHeader(c, "X-Kdbx-B64password")
	pw := Base64Decode(b64_pw)

	file := mustOpenFile(filename)
	defer file.Close()

	db = openKdbx(file, pw)
	db.UnlockProtectedEntries()
	defer db.LockProtectedEntries()

//    c.IndentedJSON(http.StatusOK, albums)
}

func postDatabase(c *gin.Context) { // update database metadata
	logs <- "Post Database"

	b64_pw := getSingleHeader(c, "X-Kdbx-B64password")
	pw := Base64Decode(b64_pw)

	file := mustOpenFile(filename)
	defer file.Close()

	db = openKdbx(file, pw)
	db.UnlockProtectedEntries()
	defer db.LockProtectedEntries()

/*
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
*/
}


/*
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// postAlbums adds an album from JSON received in the request body.
func putAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
*/
