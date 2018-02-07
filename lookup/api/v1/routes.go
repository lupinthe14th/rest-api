package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

// NewRoutes overview.
func NewRoutes() (*rest.Api, error) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/mxrecords/#host", MxRecord),
	)
	if err != nil {
		return api, err
	}
	api.SetApp(router)
	return api, nil
}
