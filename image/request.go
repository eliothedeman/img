package image

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

func ParseRequest(u *url.URL) (Request, error) {
	r := Request{}
	v := url.Values

	s = strings.Split(u.RequestURI(), "/")

	if len(s) < 1 {
		return errors.New("Image id must be given")
	} else {
		r.Id = s[0]
	}
	// default to jpeg if not found
	if v.Get("codec") != nil {
		r.Codec = v.Get("codec")
	} else {
		r.Codec = "jpg"
	}

	// defualt compression is 0
	if v.Get("compression") != nil {
		r.Compression = v.Get("compression")
	} else {
		r.Compression = "0"
	}

	// default size is small
	if v.Get("size") != nil {
		r.size = v.Get("size")
	} else {
		r.Size = "small"
	}

}
