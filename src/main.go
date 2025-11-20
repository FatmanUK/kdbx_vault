package main

import (
	"os"
	"fmt"
	"github.com/tobischo/gokeepasslib/v3"
)

func main() {
    file, _ := os.Open("examples/reading/example.kdbx")

    db := gokeepasslib.NewDatabase()
    db.Credentials = gokeepasslib.NewPasswordCredentials("abcd1234#")
    _ = gokeepasslib.NewDecoder(file).Decode(db)

    db.UnlockProtectedEntries()

    // Note: This is a simplified example and the groups and entries will depend on the specific file.
    // bound checking for the slices is recommended to avoid panics.
    entry := db.Content.Root.Groups[0].Groups[0].Entries[0]
    fmt.Println(entry.GetTitle())
    fmt.Println(entry.GetPassword())
}
