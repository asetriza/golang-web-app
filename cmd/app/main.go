package main

import (
	"github.com/asetriza/golang-web-app/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
