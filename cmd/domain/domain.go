package main

import "ports/internal/app/domain"

func main() {
	app := domain.DomainApp{}
	app.Run()
}
