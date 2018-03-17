package api

import (
	"fmt"
	"net/http"
)

var dummy = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "asfsdfd")
})
