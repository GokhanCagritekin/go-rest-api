package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GokhanCagritekin/go-rest-api/pkg/adding"
	"github.com/GokhanCagritekin/go-rest-api/pkg/listing"
	"github.com/GokhanCagritekin/go-rest-api/pkg/rest"
	"github.com/GokhanCagritekin/go-rest-api/pkg/saving"
	"github.com/GokhanCagritekin/go-rest-api/pkg/storage"
)

func main() {
	r := storage.NewStorage()
	as := adding.NewService(r)
	ls := listing.NewService(r)
	ss := saving.NewService(r)
	mux := http.NewServeMux()
	go saving.StartTask(ss)
	mux.HandleFunc("/set", rest.HandleAdding(as))
	mux.HandleFunc("/get", rest.HandleListing(ls))
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
