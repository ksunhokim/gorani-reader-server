package api

import (
	"fmt"
	"net/http"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "asfsdfd")
}
