package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

)

func TestRouter(t *testing.T)  {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
		fmt.Fprint(writer, "Hello GET")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello GET", string(body))
}