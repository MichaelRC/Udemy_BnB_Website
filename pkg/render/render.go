package render

import (
	"bytes"
	"fmt"
	"github/MRC/firstgoweb/pkg/config"
	"github/MRC/firstgoweb/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// A map of fuctions that can be used
// not built into the template 'language'
// but Go allows us to pass them in and use them this way.
var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	//eventual data to be added

	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the emplate cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreareTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//find anything that ends in .page.gohtml (* is wildcard)
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		log.Println("Error CTC 1", err)
		return myCache, err
	}

	for _, page := range pages {
		//extracts name of file
		name := filepath.Base(page)

		//creating a template set [ts]
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Error CTC 2", err)
			return myCache, err
		}

		//check to see if there are any layouts that match this template
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			log.Println("Error CTC 3", err)
			return myCache, err
		}

		//checks to see if # of matches is greater than 0
		//if TRUE parse the layout and match it to the template.
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				log.Println("Error CTC 4", err)
				return myCache, err
			}
		}

		//add template set to myCache map
		myCache[name] = ts
	}

	return myCache, nil
}
