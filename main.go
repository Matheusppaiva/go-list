package main

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/matheusppaiva/go-list/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
