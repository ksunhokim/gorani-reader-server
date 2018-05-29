package util

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write(bytes)
}
