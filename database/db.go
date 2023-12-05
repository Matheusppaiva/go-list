package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := "produtos_loja"
	host := "localhost"
	sslmode := "disable"

	conexao := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}

	// Teste a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado ao banco de dados!")

	return db
}
