package main

import (
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
)

func TestHttpResponseLayer(t *testing.T) {

	api, err := NewRoutes()
	if err != nil {
		t.Fatal(err)
	}

	handler := api.MakeHandler()

	// valid get resource
	recorded := test.RunRequest(t, handler, test.MakeSimpleRequest("GET", "http://1.2.3.4/mxrecords/ordinarius-fectum.net", nil))
	recorded.CodeIs(200)
	/*
		recorded.ContentTypeIsJson()
		recorded.BodyIs(`[
		{
			"Host": "smtp.ordinarius-fectum.net.",
			"Pref": 10
		}
		]`)
	*/

	// auto 405 on undefined route (wrong method)
	recorded = test.RunRequest(t, handler, test.MakeSimpleRequest("DELETE", "http://1.2.3.4/mxrecords/ordinarius-fectum.net", nil))
	recorded.CodeIs(405)
	/*
		recorded.ContentTypeIsJson()
		recorded.BodyIs(`{"Error":"Method not allowed"}`)
	*/

	// auto 404 on undefined route (wrong path)
	recorded = test.RunRequest(t, handler, test.MakeSimpleRequest("GET", "http://1.2.3.4/mxrecord/ordinarius-fectum.net", nil))
	recorded.CodeIs(404)
	/*
		recorded.ContentTypeIsJson()
		recorded.BodyIs(`{"Error":"Resource not found"}`)
	*/
}
