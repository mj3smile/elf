package route

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	template := new(Renderer)
	template.location = location
	template.debug = debug
	template.ReloadTemplates()
	return template
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

func Init() *echo.Echo {
	e := echo.New()

	e.Renderer = NewRenderer("templates/*.html", true)
	e.Static("/static", "static")

	e.GET("/", func(ctx echo.Context) error {
		data := M{"message": ""}
		return ctx.Render(http.StatusOK, "index.html", data)
	})

	return e
}
