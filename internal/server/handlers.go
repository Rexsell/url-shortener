package server

import (
	"net/http"
)

func ShortenerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	// read

	// validate

	// execute

	// return
}
