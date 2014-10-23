package api

import (
	"io"
	"net/http"
	"net/url"

	"github.com/eliothedeman/img/provider"
)

// RequestHandler handles an api request and responds to the requester
type RequestHandler func(w http.ResponseWriter, r Request, proto provider.Provider)

// UploadHandler handles an upload and responds to the user
type UploadHandler func(w http.ResponseWriter, r Request, hr *http.Request, proto provider.Provider)

var (
	// GET handlers for the get method
	GET map[string]RequestHandler
	// POST handlers for the post method
	POST map[string]UploadHandler
)

// init set up the handlers
func init() {
	// handlers for the get method
	GET = map[string]RequestHandler{
		"debug": DebugRequest,
		"jpg":   Download,
	}
	// handlers for the post method
	POST = map[string]UploadHandler{
		"debug": DebugUpload,
		"jpg":   Upload,
	}
}

// Router takes http requests and manages their response writers
func Router(w http.ResponseWriter, r *http.Request) {
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
	f := &provider.File{}
	f.Create(req.ID, req.Codec, req.Size, "data/")
	switch r.Method {
	case "GET":
		GET[req.Codec](w, req, f)
	case "POST":
		POST[req.Codec](w, req, r, f)
	}

}
