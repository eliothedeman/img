package api

import (
	"log"
	"net/http"
)

func DebugRequest(w http.ResponseWriter, r Request) {
	log.Println(r.Codec)
}

func DebugUpload(w http.ResponseWriter, r Request, hr *http.Request) {
	log.Println(r.Id)
}
