package main

import (
	"catering/config"
	"catering/routes"
)

func main() {
	config.Init()

	e := routes.Init()
	e.Start(":8000")

}
