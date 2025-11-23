package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func intFromString(s string) (int, error) {
	return strconv.Atoi(s)
}

func stringFromInterface(any interface{}) string {
	return fmt.Sprintf("%v", any)
}

func statusFromError(err error) int {
	// TODO: specialise error for 401 Auth, etc.
	logs <- err.Error()
	switch (err) {
		default: {
			panic(err) // to figure out what it is
		}
	}
	return http.StatusInternalServerError
}
