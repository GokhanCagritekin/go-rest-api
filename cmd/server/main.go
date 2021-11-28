package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GokhanCagritekin/go-rest-api/pkg/adding"
	"github.com/GokhanCagritekin/go-rest-api/pkg/listing"
	"github.com/GokhanCagritekin/go-rest-api/pkg/recovering"
	"github.com/GokhanCagritekin/go-rest-api/pkg/rest"
	"github.com/GokhanCagritekin/go-rest-api/pkg/saving"
	"github.com/GokhanCagritekin/go-rest-api/pkg/storage"
)

func main() {
	r := storage.NewStorage()
	addingservice := adding.NewService(r)
	listingservice := listing.NewService(r)
	savingservice := saving.NewService(r)
	recovering.Service.Recover(r)
	mux := http.NewServeMux()
	go saving.StartTask(savingservice)
	mux.HandleFunc("/set", rest.HandleAdding(addingservice))
	mux.HandleFunc("/get", rest.HandleListing(listingservice))
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
