package main

import (
	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server"
)

func main() {
	var app Server.MainRouter
	app.Router()
}
