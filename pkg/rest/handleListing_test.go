package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GokhanCagritekin/go-rest-api/pkg/listing"
)

var ls = listing.NewService(r)

func TestHandleListing(t *testing.T) {
	requestBody := map[string]string{
		"key": "test-key",
	}
	bodyData, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("TestSet Error when json.Marshal: ", err)
	}
	bufferData := bytes.NewBuffer(bodyData)

	r.Set("test-key", "exvalue")

	m := make(map[string]int)

	m["http://localhost:8080/get?key=test-key"] = 200
	m["http://localhost:8080/get?key="] = 400
	m["http://localhost:8080/get?key=notexistingkey"] = 404

	i := 0
	for k := range m {
		req, err := http.NewRequest("GET", k, bufferData)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HandleListing(ls))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != m[k] {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, m[k])
		} else {
			fmt.Printf("Req: %s status: %v want: %v \n", k, status, m[k])
		}
	}
	i++
}
