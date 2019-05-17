package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/MixinNetwork/supergroup.mixin.one/config"
	"github.com/MixinNetwork/supergroup.mixin.one/durable"
	"github.com/MixinNetwork/supergroup.mixin.one/services"
)

func main() {
	service := flag.String("service", "http", "run a service")
	flag.Parse()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DatebaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicln(err)
	}
	db.SetMaxOpenConns(128)
	db.SetMaxIdleConns(16)

	defer db.Close()
	database, err := durable.NewDatabase(context.Background(), db)
	if err != nil {
		log.Panicln(err)
	}

	switch *service {
	case "http":
		err := StartServer(database)
		if err != nil {
			log.Println(err)
		}
	default:
		go func() {
			hub := services.NewHub(database)
			err := hub.StartService(*service)
			if err != nil {
				log.Println(err)
			}
		}()

		http.ListenAndServe(fmt.Sprintf(":%d", config.HTTPListenPort+2000), http.DefaultServeMux)
	}
}
