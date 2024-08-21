package connections

import (
	"conn/internal/api/handler"
	"conn/internal/database/mehtods"
	"conn/internal/database/services"
	interface17 "conn/internal/interface"
	"database/sql"
	"log"
	_"github.com/lib/pq"
)

func NewDatabase() interface17.Books {
	db, err := sql.Open("postgres", "postgres://postgres:2005@localhost/test?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	return &mehtods.Database{Db: db}
}

func NewService() *services.Services {
	a := NewDatabase()
	return &services.Services{I: a}
}

func NewHandler() *handler.Handler {
	a := NewService()
	return &handler.Handler{S: a}
}
