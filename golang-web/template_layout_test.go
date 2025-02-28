package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml",
	))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"title": "Template Layout",
		"Name": "Kareem",
	})
}

func TestTemplateLayout(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}