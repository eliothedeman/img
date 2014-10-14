package api

import (
	"io"
	"net/http"
	"net/url"
)

type requestHandler func(w http.ResponseWriter, r Request)
type uploadHandler func(w http.ResponseWriter, r Request, hr *http.Request)

var (
	// handlers for the get method
	GET map[string]requestHandler
	// handlers for the post method
	POST map[string]requestHandler
)

// Router takes http requests and manages their response writers
func Router(w http.ResponseWriter, r *http.Request) {
	var err error
	var u *url.URL
	var req Request{}
	u, err = url.Parse(r.URL)
	if err != nil {
		io.WriteString(w, "error ya dingus\n")
		return
	}

	req, err = ParseRequest(u)
	if err != nil {
		io.WriteString(w, "error ya dingus\n")
		return
	}

	switch (r.Method)
	{
	case "GET":
		GET[req.Codec](w,req)
	case "POST":
		POST[req.Codec](w,req,r)
	}

}
