package main

import (
	adapter2 "YYY/internal/adapter"
	"YYY/internal/model"
	"YYY/internal/rest"
	"YYY/internal/service"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	conn := pg.Connect(&pg.Options{
		User:     "root",
		Password: "5432",
		Database: "postgres",
		Addr:     "db:5432",
	})

	err := createSchema(conn)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	db := service.NewService(conn)

	adapter := adapter2.NewAdapter(db)

	router := rest.Handler(adapter)
	log.Printf("Starting server on port %s.....\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.Group)(nil),
		(*model.Participant)(nil),
	}

	for _, m := range models {
		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
