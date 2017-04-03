package router

import (
	"net/http"
	"github.com/krashcan/dockertest/resources"
)

type Route struct {
	Name 		string
	Method 		string
	Pattern 	string
	HandleFunc	http.HandlerFunc
}

type Routes []Route

var routes = Routes{
		Route{
			"Reverse String.",
			"POST",
			"/reverse",
			resources.ReverseString,
		},
}
