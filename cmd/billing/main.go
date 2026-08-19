package main

import "github.com/rms-diego/store-management-eda/pkg/app"

func main() {
	app, err := app.Init("billings")
	if err != nil {
		panic(err)
	}

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
