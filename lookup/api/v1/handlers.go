package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net"
	"net/http"
)

// MxRecord returns the MX record of host passed to the parameter.
func MxRecord(w rest.ResponseWriter, req *rest.Request) {
	mx, err := net.LookupMX(req.PathParam("host"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&mx)
}
