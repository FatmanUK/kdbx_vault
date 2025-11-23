package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
//	kdbx "github.com/tobischo/gokeepasslib/v3"
)

// StatusOK = 200
// StatusNoContent = 204
// StatusUnauthorized = 401
// StatusForbidden = 403
// StatusNotFound = 404
// StatusTeapot = 418
// StatusInternalServerError = 500
// StatusBadGateway = 502
// StatusGatewayTimeout = 504
// StatusCreated = ?

func addDatabaseRoutes(r *gin.Engine, pre string) {
	r.GET(pre + "/database", getDatabase)
	r.POST(pre + "/database", postDatabase)
}

// Return database metadata.
// Marshals twice in order to delete boring data.
// Have a separate endpoint for customIcons and binaries?
func getDatabase(c *gin.Context) {
	logs <- "Get Database"
	k := Kdbx{}.Init(c)
	k, err := k.Unlock(filename)
	if err != nil {
		c.Status(statusFromError(err))
		return
	}
	defer k.Close()
	md, err := k.GetMetadata()
	if err != nil {
		c.Status(statusFromError(err))
		return
	}
	//
	k = k.Lock()
	c.JSON(http.StatusOK, md)
}

// Update database metadata.
func postDatabase(c *gin.Context) {
	logs <- "Post Database"
	k := Kdbx{}.Init(c)
	k, err := k.Unlock(filename)
	if err != nil {
		c.Status(statusFromError(err))
		return
	}
	defer k.Close()
	md, err := k.GetMetadata()
	if err != nil {
		c.Status(statusFromError(err))
		return
	}
	//
	bd, err := k.GetRequestBody()
	if err != nil {
		c.Status(statusFromError(err))
		return
	}
	logs <- string(bd)
//	logs <- string(md[0])
	c.JSON(http.StatusOK, md)

/*
	k = k.Lock()
	c.JSON(http.StatusOK, md)
*/
/*
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(statusFromError(err))
		return
	}
	var intermediate interface{}
	json.Unmarshal(reqBody, &intermediate)
	for k, v := range intermediate.(map[string]interface{}) {
		updateMetadata(db, k, stringFromInterface(v))

		//TODO
		// check we have a whatsit called k
		// check it's writable
		// write it
		logs <- (k + " -> " + v)

	}
*
	db.LockProtectedEntries()
	err = kdbx.NewEncoder(file).Encode(db)
	if err != nil {
		c.Status(statusFromError(err))
		return
	}
*/
	c.Status(http.StatusNoContent)
}
