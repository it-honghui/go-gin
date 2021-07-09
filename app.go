package main

import (
	"go-gin/config"
	"go-gin/web"
)

func init() {
	config.Setup()
}

func main() {
	web.Setup()
}
