package main

import (
	"dockertainer/api/cache"
	"dockertainer/api/databases"
)

func main() {
	Core()
}

func Core() {
	databases.Connect()
	cache.InitCache()
	go Route()
	select {}
}
