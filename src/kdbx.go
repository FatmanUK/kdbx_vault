package main

import (
	"os"
	"errors"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	kdbx "github.com/tobischo/gokeepasslib/v3"
)

type Kdbx struct {
	c *gin.Context
	db *kdbx.Database
	file *os.File
}

var auth_header string = "X-Kdbx-B64password"

func (re Kdbx) Init(c *gin.Context) Kdbx {
	re.c = c
	re.db = nil
	re.file = nil
	return re
}

func (re Kdbx) Lock() Kdbx {
	re.db.LockProtectedEntries()
	return re
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

func getSingleHeader(
		c *gin.Context,
		headerName string) (string, error) {
	header := c.Request.Header[headerName]
	l := len(header)
	m := headerName + "' headers provided"
	if l == 0 {
		return "", errors.New("Too few '" + m)
	}
	if l > 1 {
		return "", errors.New("Too many '" + m)
	}
	return header[0], nil
}

func (re Kdbx) Unlock(f string) (Kdbx, error) {
	var err error
	re.file, err = os.Open(f)
	logs <- "Opening " + f
	if err != nil {
		return re, err
	}
	b64_pw, err := getSingleHeader(re.c, auth_header)
	if err != nil {
		re.file.Close()
		return re, err
	}
	pw, err := base64Decode(b64_pw)
	if err != nil {
		re.file.Close()
		return re, err
	}
	re.db = kdbx.NewDatabase()
	re.db.Credentials = kdbx.NewPasswordCredentials(pw)
	err = kdbx.NewDecoder(re.file).Decode(re.db)
	if err != nil {
		re.file.Close()
		return re, err
	}
	re.db.UnlockProtectedEntries()
	return re, nil
}

func (re Kdbx) Close() Kdbx {
	re.file.Close()
	re.file = nil
	return re
}

func getDatabaseBoringAttrs() []string {
	return []string{"Binaries", "CustomIcons", "SettingsChanged",
		"DatabaseNameChanged", "DatabaseDescriptionChanged",
		"DefaultUserNameChanged", "MasterKeyChanged",
		"RecycleBinChanged", "EntryTemplatesGroupChanged"}
}

func getDatabaseReadOnlyAttrs() []string {
	return []string{"Generator", "Binaries", "CustomIcons",
		"SettingsChanged", "DatabaseNameChanged",
		"DatabaseDescriptionChanged",
		"DefaultUserNameChanged", "MasterKeyChanged",
		"RecycleBinChanged", "EntryTemplatesGroupChanged"}
}

func mapFromStruct(d any) (map[string]interface{}, error) {
	firstPassData, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	var intermediate interface{}
	json.Unmarshal(firstPassData, &intermediate)
	finalData := intermediate.(map[string]interface{})
	return finalData, nil
}

func (re Kdbx) GetMetadata() (map[string]interface{}, error) {
	finalData, err := mapFromStruct(re.db.Content.Meta)
	if err != nil {
		return nil, err
	}
	boringData := getDatabaseBoringAttrs()
	for _, v := range boringData {
		delete(finalData, v)
	}
	return finalData, nil
}

func (re Kdbx) GetRequestBody() ([]byte, error) {
	reqBody, err := ioutil.ReadAll(re.c.Request.Body)
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}
