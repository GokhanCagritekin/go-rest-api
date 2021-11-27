package rest

import (
	"net/http"

	"github.com/GokhanCagritekin/go-rest-api/pkg/adding"
)

func HandleAdding(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in path", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val := r.URL.Query().Get("value")
		as.Set(key, string(val))
		w.WriteHeader(http.StatusOK)
	}
}
