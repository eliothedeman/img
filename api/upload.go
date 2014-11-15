package api

import (
	"log"
	"net/http"

	"github.com/eliothedeman/img/provider"
)

// Upload handles an upload api request
func Upload(w http.ResponseWriter, r Request, hr *http.Request, proto provider.Provider) {
	n := BUFFERSIZE
	buff := make([]byte, BUFFERSIZE)
	off := int64(0)
	done := false
	var err error

	for !done {
		n, err = hr.Body.Read(buff)
		if err != nil {
			if err.Error() == "EOF" {
				done = true
			} else {
				http.Error(w, err.Error(), 500)
				return
			}

		}
		n, err = proto.WriteAt(buff[0:n], off)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		off += int64(n)
	}
	log.Printf("%d bytes written to %s", off, proto.Location())
}
