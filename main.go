package main

import (
	"github.com/adrianorocha-dev/html-deez-nuts/model"
	"github.com/adrianorocha-dev/html-deez-nuts/routes"
)

func main() {
	model.Setup()
	routes.SetupAndRun()
}
