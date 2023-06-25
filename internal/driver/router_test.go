package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterWorker(t *testing.T) {
	values := map[string]string{"url": fmt.Sprintf("http://127.0.0.1:%s", "8081")}
	jsonData, _ := json.Marshal(values)
	req := httptest.NewRequest(
		"POST",
		"/reduce",
		bytes.NewBuffer(jsonData),
	)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	registerHandle(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.JSONEq(
		t,
		`{"message":"reduceTask called"}`,
		string(resBody))
}
