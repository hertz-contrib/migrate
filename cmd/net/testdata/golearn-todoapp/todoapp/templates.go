package todoapp

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Looks for templates and preload (parse) them in memory.
func PreLoadTemplates() {
	templatesDir := AppConfig.TemplateAbsPath

	layouts, err := filepath.Glob(filepath.Join(templatesDir, "layouts", "*.html"))
	if err != nil {
		log.Fatal(err)
	}

	includes, err := filepath.Glob(filepath.Join(templatesDir, "includes", "*.html"))
	if err != nil {
		log.Fatal(err)
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		AppConfig.Templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}
}

// Render a given template (that was pre-loaded).
func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	// Ensure the template exists in the map.
	tpl, ok := AppConfig.Templates[name]
	if !ok {
		return fmt.Errorf("the template %s does not exist", name)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tpl.ExecuteTemplate(w, name, data)
}
