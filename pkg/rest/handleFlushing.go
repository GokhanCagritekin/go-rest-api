package rest

import (
	"fmt"
	"net/http"

	"github.com/GokhanCagritekin/go-rest-api/pkg/flushing"
)

func HandleFlushing(fs flushing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fs.DeleteAll()
		if err != nil {
			fmt.Println("Error occured while Flushing")
		}
		w.WriteHeader(http.StatusOK)
	}
}
