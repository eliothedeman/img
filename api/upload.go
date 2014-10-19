package api

import (
	"io/ioutil"
	"net/http"

	"github.com/eliothedeman/img/provider"
)

// Upload handles an upload api request
func Upload(w http.ResponseWriter, r Request, hr *http.Request, proto provider.Provider) {
	b, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		http.Error(w, err.Error(), 1000)
	}
	_, err = proto.Write(b)
	if err != nil {
		http.Error(w, err.Error(), 1001)
	}
}
