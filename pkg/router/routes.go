package router

import "net/http"

func (a *API) download(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
	}

}

func (a *API) upload(w http.ResponseWriter, req *http.Request) {

}

func (a *API) search(w http.ResponseWriter, req *http.Request) {
	req.URL.Query()

}
