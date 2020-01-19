package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/jcfug8/ai-writer/services/persist/migrations"
	"github.com/jcfug8/ai-writer/services/persist/service"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var addr string
var dbAddr string
var dbUser string
var dbPass string
var dbDatabase string

func init() {
	flag.StringVar(&addr, "a", "127.0.0.1:50051", "Address Persist service will be served at")
	flag.StringVar(&dbAddr, "db", "127.0.0.1:3306", "Address Database can be reach at")
	flag.StringVar(&dbUser, "du", "ai_writer", "Database user")
	flag.StringVar(&dbPass, "dp", "ai_writer", "Database password")
	flag.StringVar(&dbDatabase, "dn", "ai_writer", "Database name")
}

func main() {
	flag.Parse()
	log.WithFields(log.Fields{
		"a":  addr,
		"db": dbAddr,
		"du": dbUser,
		"dp": dbPass,
		"dn": dbDatabase,
	}).Info("parsed flags")

	persist := service.NewService(&service.Opts{
		Addr: addr,
	})

	db, err := CreateDatabaseConnection(dbAddr, dbUser, dbPass, dbDatabase)
	if err != nil {
		log.Panic("could not create database connection")
	}
	migrations.Run(db)

	persist.RegisterDatabase(db)

	log.Info("about to serve persist")
	if err := persist.Serve(); err != nil {
		log.Errorf("control http server closed in error: %s", err)
	}
}

func CreateDatabaseConnection(addr, user, pass, name string) (*sql.DB, error) {
	end := time.Now().Add(time.Minute)

	for time.Now().Before(end) {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, addr, name))
		// if there is an error opening the connection, handle it
		if err != nil {
			log.Warning("db connection could not be made")
			time.Sleep(time.Second)
			continue
		}
		log.Info("database connection made - attemping to ping")
		for time.Now().Before(end) {
			if db.Ping() != nil {
				log.Warning("db ping failed")
				time.Sleep(time.Second)
				continue
			}
			log.Info("database connection ready")
			return db, nil
		}
	}
	return nil, errors.New("Unable to create database connection")
}
