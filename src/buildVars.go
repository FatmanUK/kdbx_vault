package main

// These are set during build.
var BUILD_MODE string
var APP_NAME string
var VERSION string
var KV_PORT string

// These are not.
var logs chan string
var filename string
