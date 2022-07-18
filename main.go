package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

type ServiceInfo struct {
	Title string
}

var (
	tr = &Renderer{templates: template.Must(template.ParseGlob("views/*.html"))}
)

var serviceInfo = ServiceInfo{
	"Valorant Agent",
}

func main() {
	e := echo.New()

	e.Renderer = tr

	e.GET("/", example)

	e.GET("/home", home)

	e.Logger.Fatal(e.Start(":1323"))
}

// TODO: 後で消す
func example(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func home(c echo.Context) error {
	data := struct {
		ServiceInfo
		Content string
	}{
		ServiceInfo: serviceInfo,
		Content:     "Jett",
	}

	return c.Render(http.StatusOK, "page1", data)
}
