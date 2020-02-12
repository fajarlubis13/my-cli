package main

import (
	"{{ toDelimeted .ProjectName 45 }}/routes"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	routes.NewRoutes().Run()
}
