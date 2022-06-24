package index

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	response := &IndexResponse{
		Status: 200,
		Message: "Welcome to art.jackli.dev API",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}