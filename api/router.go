package api

import (
	"io"
	"net/http"
	"net/url"

	"github.com/eliothedeman/img/image"
)

// Router takes http requests and manages their response writers
func Router(w http.ResponseWriter, r *http.Request) {
	req := image.Request{}

	u, err := url.Parse(r.URL)
	if err != nil {
		io.WriteString(w, "error ya dingus\n")
		return
	}
}
