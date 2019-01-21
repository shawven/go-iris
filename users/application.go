package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	registerControllers(app)

	_ = app.Run(
		iris.Addr(":8080"),
		iris.WithOptimizations,
		iris.WithCharset("utf-8"),
	)
}

func registerControllers(app *iris.Application) {

}
