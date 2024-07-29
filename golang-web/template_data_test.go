package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TemplateDataMap(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name": "Kareem",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada",
		},
	})
}

func TestTemplateDataMap(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Page struct{
	Title string
	Name string
	Address Address
}
type Address struct{
	Street string
}

func TemplateStruct(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Map",
		Name: "Kareem",
		Address: Address{
			Street: "Jalan Belum Ada",
		},
	})
}

func TestTemplateStruct(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateStruct(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}