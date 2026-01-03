package main

import "net/http"

func server() {
	ServeMux := http.NewServeMux()
	Server := &http.Server{
		Handler: ServeMux,
		Addr:    ":8080",
	}
	Server.ListenAndServe()
}
