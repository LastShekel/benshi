package driver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// parseRequest into json based map
func parseRequest(r *http.Request) map[string]string {
	data := make(map[string]string)
	var p []byte
	for {
		p1 := make([]byte, 1)
		_, err := io.ReadFull(r.Body, p1)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil
		}
		p = append(p, p1...)
	}

	json.Unmarshal(p, &data)
	return data
}

// registerHandle handles all worker register requests
func registerHandle(w http.ResponseWriter, r *http.Request) {
	data := parseRequest(r)
	controller.regMutex.Lock()
	_, ex := controller.registeredWorkers[data["url"]]
	if !ex {
		controller.registeredWorkers[data["url"]] = true
		log.Printf("Worker registered: %s\n", data["url"])
	}
	controller.regMutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte(`{"message": "Ok"}`))
	if err != nil {
		return
	}
}

// doneHandle handles all worker done requests
func doneHandle(w http.ResponseWriter, r *http.Request) {
	//data := parseRequest(r)
	controller.doneMutex.Lock()
	url := parseRequest(r)["url"]
	delete(controller.sentTasks, url)
	log.Printf("Worker done: %s\n", url)
	controller.workersQueue <- url
	controller.doneMutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Ok"}`))
}

var controller *TaskController

// NewRouter creates router api for driver
func NewRouter(c *TaskController) http.Handler {
	controller = c
	api := mux.NewRouter()
	api.HandleFunc("/register", registerHandle).Methods(http.MethodPost)
	api.HandleFunc("/done", doneHandle).Methods(http.MethodPost)
	return api
}
