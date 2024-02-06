package todoapp

import "html/template"

// The application configuration
type Configuration struct {
	Hostname        string
	Port            uint
	AppRootPath     string
	TemplateAbsPath string
	Templates       map[string]*template.Template
}

// Create the default configuration.
// This will get populated by any user passed data.
var AppConfig = &Configuration{
	"127.0.0.1",
	8000,
	"",
	"",
	make(map[string]*template.Template)}
