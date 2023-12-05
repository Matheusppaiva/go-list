package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/matheusppaiva/go-list/routes"
)

func main() {
	routes.CarregaRotas()

	matheus := os.Getenv("MATHEUS")
	// log.Default(matheus)
	fmt.Println(matheus)

	http.ListenAndServe(":8000", nil)

}
