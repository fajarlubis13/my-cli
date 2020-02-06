package main {{ toLowerCamel .Title }}

import (
	"hk-jadwal-teknik/routes"
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
