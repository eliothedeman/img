package api

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type RequestHandler func(w http.ResponseWriter, r Request)
type UploadHandler func(w http.ResponseWriter, r Request, hr *http.Request)

var (
	// handlers for the get method
	GET map[string]RequestHandler
	// handlers for the post method
	POST map[string]UploadHandler
)

func init() {
	// handlers for the get method
	GET = map[string]RequestHandler{
		"jpg": DebugRequest,
	}
	// handlers for the post method
	POST = map[string]UploadHandler{
		"jpg": DebugUpload,
	}
}

// Router takes http requests and manages their response writers
func Router(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.String())
	var err error
	var u *url.URL
	var req Request
	u = r.URL
	if err != nil {
		io.WriteString(w, "error ya dingus\n")
		return
	}

	req, err = ParseRequest(u)
	if err != nil {
		io.WriteString(w, "error ya dingus\n")
		return
	}

	switch r.Method {
	case "GET":
		GET[req.Codec](w, req)
	case "POST":
		POST[req.Codec](w, req, r)
	}

}
