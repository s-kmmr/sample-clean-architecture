package main

import "github.com/s-kmmr/sample-clean-architecture/infrastructure/router"

func main() {
	r := router.NewRouter()
	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
