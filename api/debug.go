package api

import (
	"log"
	"net/http"

	"github.com/eliothedeman/img/provider"
)

// DebugRequest debug RequestHandler
func DebugRequest(w http.ResponseWriter, r Request) {
	log.Println(r.Codec)
}

// DebugUpload debug UploadHandler
func DebugUpload(w http.ResponseWriter, r Request, hr *http.Request, proto provider.Provider) {
	log.Println(r.ID)
}
