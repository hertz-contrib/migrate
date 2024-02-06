package routes

import (
	"github.com/hertz-contrib/migrate/cmd/net/testdata/GoWeb/app"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/GoWeb/controllers"
	"github.com/hertz-contrib/migrate/cmd/net/testdata/GoWeb/middleware"
	"net/http"
)

// Post defines all project post routes
func Post(app *app.App) {
	// Post controller struct initialize
	postController := controllers.Post{
		App: app,
	}

	// User authentication
	http.HandleFunc("/register-handle", middleware.Csrf(postController.Register))
	http.HandleFunc("/login-handle", middleware.Csrf(postController.Login))
}
