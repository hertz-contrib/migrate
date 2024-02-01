package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(
		server.WithWriteTimeout(30),
	)
	h.GET("/", _sayhelloName)
	h.Spin()
}

func _sayhelloName(ctx context.Context, c *app.RequestContext) {
	//uri := c.Request.URI().String()
	//host := string(c.Host())
	//println(uri)
	//m := string(c.Method())
	//c.SetBodyString("Hello afei!")
}
