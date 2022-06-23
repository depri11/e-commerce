package main

import (
	"log"

	"github.com/depri11/e-commerce/src/routers"
)

func main() {
	e, err := routers.SetupRouters()
	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":4000"))

}
