package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	_ "embed"
)

// golang template
func SimpleHTML(writer http.ResponseWriter, request *http.Request)  {
	templatText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templatText)
	// if err != nil{
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templatText))

	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func TestTemplateGolang(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


// golang template file
func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateFileGolang(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


// golang template Directory
func SimpleHTMLDirectory(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateDirectoryGolang(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLDirectory(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


//go:embed templates/*.gohtml
var templates embed.FS
func SimpleHTMLEmbed(writer http.ResponseWriter, request *http.Request)  {

	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateEmbedGolang(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLEmbed(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
