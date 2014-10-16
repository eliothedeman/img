package api

import (
	"errors"
	"net/url"
	"strings"
)

// Request describes a image request
type Request struct {
	Codec       string
	Compression string
	Id          string
	Size        string
}

// ParseRequest create a Request from a url
func ParseRequest(u *url.URL) (Request, error) {
	r := Request{}
	v := u.Query()

	s := strings.Split(u.RequestURI(), "/")

	if len(s) < 1 {
		return Request{}, errors.New("Image id must be given")
	} else {
		r.Id = s[0]
	}
	// default to jpeg if not found
	if v.Get("codec") != "" {
		r.Codec = strings.ToLower(v.Get("codec"))
	} else {
		r.Codec = "jpg"
	}

	// defualt compression is 0
	if v.Get("compression") != "" {
		r.Compression = strings.ToLower(v.Get("compression"))
	} else {
		r.Compression = "0"
	}

	// default size is small
	if v.Get("size") != "" {
		r.Size = strings.ToLower(v.Get("size"))
	} else {
		r.Size = "small"
	}
	return r, nil

}
