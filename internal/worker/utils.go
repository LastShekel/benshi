package worker

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func createFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func SendJson(url string, values map[string]string) (resp *http.Response, err error) {
	jsonData, _ := json.Marshal(values)
	return http.Post(url, "application/json", bytes.NewBuffer(jsonData))
}
