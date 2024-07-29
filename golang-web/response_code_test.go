package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func ResponseCode(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest) // Bad Request
		fmt.Fprint(writer, "name is empty")
	}else{
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestResoponseCodeInvalid(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}

func TestResoponseCodeValid(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=kareem", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}