package controller

import "net/http"

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/API", create())
	mux.HandleFunc("/", tracker())
	return mux
}

