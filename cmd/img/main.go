package main

import "github.com/eliothedeman/img/api"

func main() {
	s := api.Server{}
	s.Serve("8080")
}
