package main

import (
	"github.com/depri11/e-commerce/src/routers"
)

func main() {
	e, err := routers.SetupRouters()
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":4000"))

}
