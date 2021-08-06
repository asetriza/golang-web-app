package main

import (
	"log"

	"github.com/asetriza/golang-web-app/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
