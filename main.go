package main

import (
	"assignment3/routers"
)

func main() {
	const PORT = "localhost:8080"

	routers.StartServer().Run(PORT)
}
