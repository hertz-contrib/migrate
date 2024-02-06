package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hertz-contrib/migrate/cmd/net/testdata/golearn-todoapp/todoapp"
)

// Resolves the different paths used by the server application
//   - The root app (current working directory by default);
//   - The HTML templates path (“{app_root} + /templates“).
func resolvePaths(config *todoapp.Configuration) {
	var err error

	// Retrieve the root absolute path (main.go's directory)
	config.AppRootPath, err = filepath.Abs(".")

	// If there was an error, give up
	if err != nil {
		log.Fatal(err)
	}
	// Resolve the templates path ("{root}/templates")
	config.TemplateAbsPath = filepath.Join(config.AppRootPath, "templates")
}

// Setup the go-flags and parse the user passed arguments,
// returns the generated application configuration file.
func parseArguments() {
	config := todoapp.AppConfig

	// The application's hostname to listen on
	flag.StringVar(
		&config.Hostname, "addr",
		config.Hostname, "The address to listen on.")

	// FIXME: we want a uint16 here, but is not existing. Create one?
	// The port to listen on
	flag.UintVar(
		&config.Port, "port",
		config.Port, "The port to listen to.")

	// Parse the arguments and populate the default configuration.
	flag.Parse()
}

// Start the application's server an listen on a given point
// (in accordance to the configuration file).
// Can also take a non-default HTTP server handler to use
// (can be nil to use default).
func serve(config *todoapp.Configuration, handler http.Handler) {
	listenAddr := fmt.Sprintf("%s:%d", config.Hostname, config.Port)

	log.Printf("Starting http://%s...", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, handler))
}

// “todoapp“'s entry point. Setup the application
// and listening for HTTP requests forever.
func main() {
	httpRouter := chi.NewRouter()
	config := todoapp.AppConfig

	parseArguments()
	resolvePaths(config)

	todoapp.PreLoadTemplates()
	todoapp.SetupRoutes(httpRouter)

	serve(config, httpRouter)
}
