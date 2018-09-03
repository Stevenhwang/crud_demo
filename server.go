package main

import (
	"awesomeProject/router"
)

func main() {
	// set router
	router.SetRouter()
	// start server
	router.Server.Logger.Fatal(router.Server.Start(":1323"))
}
