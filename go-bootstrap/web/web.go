package main

import (
	"web/src/router"
	"web/src/sess"
)

func main() {

	//db connect
	//db.Init()

	e := router.Init()

	middle := sess.Init()
	e.Use(middle)
	e.Use(sess.Handler)

	e.Logger.Fatal(e.Start(":1323"))
}
