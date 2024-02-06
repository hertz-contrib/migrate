package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net/http"
)

func main() {
	h := server.Default(
		server.WithWriteTimeout(30),
	)
	h.GET("/", _sayhelloName)
	h.Spin()
}

func _sayhelloName(ctx context.Context, c *app.RequestContext) {
	c.Redirect(http.StatusMovedPermanently, []byte("/api/v1/healthz"))
	//switch string(c.Method()) {
	//}
	//if string(c.Method()) == "POST" {
	//}
	//form, err := c.MultipartForm()
	//value := form.Value
	//value := string(c.FormValue("name"))
	//uri := c.Request.URI().String()
	//host := string(c.Host())
	//println(uri)
	//m := string(c.Method())
	//c.SetBodyString("Hello afei!")
	//s := c.URI().String()
	//file, err := c.Request.FormFile("s")
}
