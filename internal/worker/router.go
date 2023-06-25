package worker

import (
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

var intermediatePath = "intermediate"
var outPath = "out"
var driverUrl = "http://127.0.0.1"
var workerUrl = "http://127.0.0.1"

// NewRouter creates router api for router
func NewRouter(c Conf) http.Handler {
	intermediatePath = path.Join(c.Files, intermediatePath)
	outPath = path.Join(c.Files, outPath)
	driverUrl = "http://127.0.0.1:" + c.DriverPort
	workerUrl = "http://127.0.0.1:" + c.WorkerPort

	createFolder(intermediatePath)
	createFolder(outPath)
	api := mux.NewRouter()
	api.HandleFunc("/map", mapHandle).Methods(http.MethodPost)
	api.HandleFunc("/reduce", reduceHandle).Methods(http.MethodPost)
	return api
}
