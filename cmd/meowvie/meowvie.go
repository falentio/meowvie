package main

import (
	"meowvie/internal"
)

func main() {
	app := internal.NewApplication()

	if err := app.Listen(); err != nil {
		panic(err)
	}
}
