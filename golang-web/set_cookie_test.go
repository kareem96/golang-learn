package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request)  {
	cookie := new(http.Cookie)
	cookie.Name = "Abdul-Karim"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success Create Cookie")
	
}

func GetCookie(writer http.ResponseWriter, request *http.Request)  {
	cookie, err := request.Cookie("Abdul-Karim")
	if err != nil{
		fmt.Fprint(writer, "No Cookie")
	}else{
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}

}

func TestSetCookie(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Kareem", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies{
		fmt.Printf("Cookie %s %s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "Abdul-Karim"
	cookie.Value = "Kareem"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}