package routes

import (
	"fmt"
	"net/http"
)

func RegisterAPI(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message":"hello mom"}`)
	})
}
