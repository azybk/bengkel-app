package main

import (
	"bengkel_app/internal/component"
	"bengkel_app/internal/config"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetDatabaseConnection(conf)
}