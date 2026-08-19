package main

import "github.com/rms-diego/store-management-eda/pkg/app"

func main() {
	app, _, err := app.Init()
	if err != nil {
		panic(err)
	}

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
