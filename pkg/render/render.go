package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/madegun/bookings/models"
	"github.com/madegun/bookings/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//AddDefaultData handler for model data
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CRSFToken = nosurf.Token(r)

	return td
}

//NewTemplate set the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

//RenderTemplate using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, html string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseChace {
		tc = app.TemplateChace
	} else {
		tc, _ = CreateTemplateChace()

	}
	t, ok := tc[html]
	if !ok {
		log.Fatal("tidak bisa dapat template dari templates chace")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error write template ke browser", err)
	}

}

//CreateTemplateChace
func CreateTemplateChace() (map[string]*template.Template, error) {

	myChace := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myChace, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myChace, err
		}

		cocok, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myChace, err
		}

		if len(cocok) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myChace, err
			}
		}

		myChace[name] = ts
	}
	return myChace, nil

}
