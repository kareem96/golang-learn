package golangweb

import (
	"fmt"
	"net/http"
	"testing"
	_ "embed"
)

func ServeFile(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Query().Get("name") != ""{
		http.ServeFile(writer, request, "./resources/ok.html")
	}else{
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotfound string

func ServeFileGolangEmbed(writer http.ResponseWriter, request *http.Request)  {
	if request.URL.Query().Get("name") != ""{
		fmt.Fprint(writer, resourcesOk)
	}else{
		fmt.Fprint(writer, resourcesNotfound)
	}
}

func TestServeFileGolangEmbed(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}