package main

import (
	"go-practice/handlers"
)

func main() {
	r := handlers.SetUpRoutes()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
