//go:build testintegration

package api

import (
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// FooBarRequest represents body of FooBar request.
type FooBarRequest struct {
	// Foo param
	Foo string `json:"foo"`
	// Bar params
	Bar []int `json:"bar"`
	// User param
	User User `json:"user"`
}

// FooBarResponse represents body of FooBar response.
type FooBarResponse struct {
	Baz struct {
		Prop string `json:"prop"`
	} `json:"baz"`
}

// FooBarHandler handles incoming foobar requests
func FooBarHandler(w http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func doSthWithRequest(req FooBarRequest) FooBarResponse {
	_ = "STUB: not implemented"
	return *new(FooBarResponse)
}
