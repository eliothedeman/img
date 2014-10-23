package main

import (
	"log"

	"github.com/eliothedeman/img/api"
)

func main() {
	log.SetFlags(log.Llongfile)
	s := api.Server{}
	s.Serve("8080")
}
