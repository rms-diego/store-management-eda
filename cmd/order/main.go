package main

import (
	"github.com/rms-diego/store-management-eda/pkg/app"
)

func main() {
	a, _, err := app.Init()
	if err != nil {
		panic(err)
	}

	if err := a.ListenAndServe(); err != nil {
		panic(err)
	}
}
