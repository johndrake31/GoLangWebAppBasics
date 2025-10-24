package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateDir = "./templates"

// RenderTemplate renders a template using the html/template package
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	path := filepath.Join(templateDir, tmpl)
	
	t, err := template.ParseFiles(path)
	
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
