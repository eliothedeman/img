package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eliothedeman/img/provider"
)

// Download send a data to the user using the data provider
func Download(w http.ResponseWriter, r Request, proto provider.Provider) {
	defer func() {
		err := proto.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	b := make([]byte, BUFFERSIZE)
	var err error
	var off int
	n := BUFFERSIZE
	for n == BUFFERSIZE {
		// read bytes from buffer
		n, err = proto.ReadAt(b, int64(off))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		off += n
		// write bytes to requester
		_, err = w.Write(b)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Printf("Wrote %d bytes to client\n", n)
	}
}
