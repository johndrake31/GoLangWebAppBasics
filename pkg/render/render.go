package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateDir = "./templates"

// RenderTemplate renders a template using the html/template package
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	path := filepath.Join(templateDir, tmpl)

	t, err := template.ParseFiles(path, filepath.Join(templateDir, "base.layout.tmpl"))

	// Handle error if template parsing fails
	if err != nil {
		log.Printf("ParseFiles(%s) failed: %v", path, err)
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	// Execute the template with no dynamic data
	if err := t.Execute(w, nil); err != nil {
		log.Printf("Execute(%s) failed: %v", tmpl, err)
		http.Error(w, "template execute error", http.StatusInternalServerError)
		return
	}
}

var tc = make(map[string]*template.Template)

// RenderTemplateTest renders a template using a template cache
func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		// need to create the template if not found
	} else {
		// we have it in the map
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
}

// CreateTemplateCache creates a template cache
func CreateTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("%s/%s", templateDir, t),
		fmt.Sprintf("%s/%s", templateDir, "base.layout.tmpl"),
	}

	// parse the template files
	tmpl, err := template.ParseFiles(templates...)
	// Handle error if template parsing fails
	if err != nil {
		log.Printf("ParseFiles(%s) failed: %v", tmpl, err)
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}

	return nil
}
