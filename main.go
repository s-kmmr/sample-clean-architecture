package main

import (
	_ "github.com/s-kmmr/sample-clean-architecture/docs"

	"github.com/s-kmmr/sample-clean-architecture/infrastructure/injector"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/router"
)

// @title sample-clean-architecture API
// @version 1.0
// @license.name kmmr
// @description sample-clean-architecture OpenAPI
// @termsOfService none
func main() {
	i := injector.NewInjector()
	router.NewRouter(i).Start()
}
