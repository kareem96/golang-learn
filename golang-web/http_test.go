package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func HelloHandler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(writer, "Hello World!")
}
func TestHtttp(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}

// Query Parameter
func SayHello(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name == ""{
		fmt.Fprint(writer, "Hello")
	}else{
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Kareem", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}

// Query Parameter Multiple
func QueryPamaterMultiple(writer http.ResponseWriter, request *http.Request)  {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestQueryParameterMultiple(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Kareem&last_name=Karim", nil)
	recorder := httptest.NewRecorder()

	QueryPamaterMultiple(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}

// Query Parameter Multiple Values
func QueryPamaterMultipleValues(writer http.ResponseWriter, request *http.Request)  {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestQueryParameterMultiplValues(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Abdul&name=Karim&name=Melayu", nil)
	recorder := httptest.NewRecorder()

	QueryPamaterMultipleValues(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	bodyString := string(body)
	fmt.Println(bodyString)
}