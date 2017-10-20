package wallet

import (
	"net/http"
)
// Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{
	/*
	Route{
		"GetWallet",                                     // Name
		"GET",                                            // HTTP method
		"/v1/accounts/{walletId}",                          // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			w_id := vars["walletId"]
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"+w_id))
		},
	},*/
	Route{
		"GetWallet",                                     // Name
		"GET",                                            // HTTP method
		"/v1/accounts/{walletId}",                          // Route pattern
		getwallet,
	},
}