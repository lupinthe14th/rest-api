package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net"
	"net/http"
)

func MxRecord(w rest.ResponseWriter, req *rest.Request) {
	mx, err := net.LookupMX(req.PathParam("host"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&mx)
}
