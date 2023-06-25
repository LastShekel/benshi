package worker

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestPostMap(t *testing.T) {
	router := NewRouter(Conf{})

	req, _ := http.NewRequest("POST", "map", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	fmt.Println(rr.Code, rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}

func TestMapTask(t *testing.T) {
	req := getMapRequest("0")
	rr := httptest.NewRecorder()

	mapHandle(rr, req)
	req = getMapRequest("1")
	mapHandle(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.JSONEq(
		t,
		`{"message":"mapHandle called"}{"message":"mapHandle called"}`,
		string(resBody))
}
func getMapRequest(taskId string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for i, filename := range []string{
		"./testdata/pg-metamorphosis.txt",
		"./testdata/pg-sherlock_holmes.txt",
		"./testdata/pg-tom_sawyer.txt",
	} {
		part, _ := writer.CreateFormFile("file"+strconv.Itoa(i), filename)
		data, _ := os.ReadFile(filename)
		_, err := part.Write(data)
		if err != nil {
			return nil
		} // <-- content is the []byte
	}

	err := writer.Close()
	if err != nil {
		return nil
	}
	req := httptest.NewRequest(
		"POST",
		"/map",
		body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("M", "2")
	req.Header.Add("taskId", taskId)
	return req
}
