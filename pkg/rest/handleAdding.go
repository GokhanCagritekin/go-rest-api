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
		val := r.URL.Query().Get("value")
		if val == "" {
			http.Error(w, "missing value name in path", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		as.Set(key, val)
		w.WriteHeader(http.StatusOK)
	}
}
