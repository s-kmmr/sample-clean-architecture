package main

import (
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/injector"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/router"
)

func main() {
	i := injector.NewInjector()
	router.NewRouter(i).Start()
}
