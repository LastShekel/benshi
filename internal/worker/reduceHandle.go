package worker

import (
	"context"
	"log"
	"net/http"
	"os"
)

// reduceHandle is handle for router. It takes extra data from headers. it runs reduce file task
func reduceHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	entries, err := os.ReadDir(intermediatePath)
	if err != nil {
		log.Fatal(err)
	}
	if len(entries) == 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message": "Intermediate folder is empty, perhaps there wasn't any map task run"}`))
		return
	}
	id := r.Header["Taskid"][0]
	ctx := context.Background()
	ctx = context.WithValue(ctx, "taskId", id)
	ctx = context.WithValue(ctx, "folder", intermediatePath)
	ctx = context.WithValue(ctx, "out", outPath)
	ctx = context.WithValue(ctx, "driverUrl", driverUrl)
	ctx = context.WithValue(ctx, "workerUrl", workerUrl)
	go reduceFiles(ctx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "reduceHandle called"}`))
}
