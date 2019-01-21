package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.Default()

	app.Handle("Get", "/", func(ctx context.Context) {
		_, _ = ctx.HTML("hello")
	})

	app.Get("/ping", func(ctx context.Context) {
		_, _ = ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx context.Context) {
		_, _ = ctx.JSON(iris.Map{"message": "hello world form iris"})
	})

	_ = app.Run(iris.Addr("localhost:8080"))
}
