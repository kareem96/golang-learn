package golangweb

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templatesAutoEscapte embed.FS

var myTemplatesAutoEscape = template.Must(template.ParseFS(templatesAutoEscapte, "templates/*.gohtml"))

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request)  {
	myTemplatesAutoEscape.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body": "<p>Selamat Belajar Go-Lang Web<script>alert('Anda di hack')</script></p>",
	})
}
func TestTemplateAutoEscape(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TestTemplateAutoEscapeServer(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateDisableAutoEscape(writer http.ResponseWriter, request *http.Request)  {
	myTemplatesAutoEscape.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body": template.HTML("<h1>Selamat Belajar Go-Lang Web<script>alert('Anda di hack')</script></h1>"),
	})
}
func TestTemplateDisableAutoEscape(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDisableAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TestTemplateDisableAutoEscapeServer(t *testing.T)  {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateDisableAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}