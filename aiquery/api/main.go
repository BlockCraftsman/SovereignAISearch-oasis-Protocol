package main

import (
	"log"

	"aiquery/api/handler"
)

func main() {
	initAll()
	g.POST("/query", handler.Search)
	startGin()
}

func init() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}
