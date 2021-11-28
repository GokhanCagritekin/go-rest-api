package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GokhanCagritekin/go-rest-api/pkg/adding"
	"github.com/GokhanCagritekin/go-rest-api/pkg/storage"
)

var r = storage.NewStorage()
var as = adding.NewService(r)

func TestHandleAdding(t *testing.T) {
	requestBody := map[string]string{
		"key":   "test-key",
		"value": "test-value",
	}
	bodyData, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("TestSet Error when json.Marshal: ", err)
	}
	bufferData := bytes.NewBuffer(bodyData)

	m := make(map[string]int)

	m["http://localhost:8080/set?key=gok&value=ncv"] = 200
	m["http://localhost:8080/set?key=&value=ncv"] = 400
	m["http://localhost:8080/set?key=g&value="] = 400

	i := 0
	for k := range m {
		req, err := http.NewRequest("POST", k, bufferData)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HandleAdding(as))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != m[k] {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		} else {
			fmt.Printf("Req: %s status: %v want: %v \n", k, status, m[k])
		}
	}
	i++
}
