package worker

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

// mapHandle is handle for router. It takes extra data from headers and files from body.
func mapHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	M, err := strconv.Atoi(r.Header["M"][0])
	if err != nil {
		log.Printf("Warning failed to parse %s setting M to 1\n", r.Header["M"][0])
		M = 1
	}
	id := r.Header["Taskid"][0]
	contentType := r.Header.Get("Content-Type")
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "Problem with Content-Type header"}`))
		return
	}
	log.Printf("Recieved map task id=%s", id)
	files := readMultipart(r, mediaType, params)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "M", M)
	ctx = context.WithValue(ctx, "files", files)
	ctx = context.WithValue(ctx, "taskId", id)
	ctx = context.WithValue(ctx, "folder", intermediatePath)
	ctx = context.WithValue(ctx, "driverUrl", driverUrl)
	ctx = context.WithValue(ctx, "workerUrl", workerUrl)
	go mapFiles(ctx)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "mapHandle called"}`))
}

// readMultipart reading data from request into files.
func readMultipart(r *http.Request, mediaType string, params map[string]string) []string {
	var files []string
	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			// p.FormName() is the name of the element.
			// p.FileName() is the name of the file (if it's a file)
			// p is an io.Reader on the part

			// The following code prints the part for demonstration purposes.
			data, err := ioutil.ReadAll(p)
			if err != nil {
				log.Fatal(err)
			}
			files = append(files, string(data))
		}
	}
	return files
}
