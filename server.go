package main

import (
	"crud_demo/router"
)

func main() {
	// set router
	router.SetRouter()
	// start server
	router.Server.Logger.Fatal(router.Server.Start(":1323"))
}
