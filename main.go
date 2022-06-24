package main

import (
	"log"
	"os"

	"github.com/depri11/e-commerce/src/routers"
)

func main() {
	e, err := routers.SetupRouters()
	if err != nil {
		log.Fatal(err)
	}

	var addrs string = "0.0.0.0:4000"

	if pr := os.Getenv("PORT"); pr != "" {
		addrs = "0.0.0.0:" + pr
	}

	e.Logger.Fatal(e.Start(addrs))

}
