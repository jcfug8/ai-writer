package migrations

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
	log "github.com/sirupsen/logrus"
)

var (
	migrations = []darwin.Migration{
		{
			Version:     1,
			Description: "Creating table users",
			Script: `CREATE TABLE users (
						id INT 		auto_increment, 
						email 		VARCHAR(255) NOT NULL UNIQUE,
						hashed_password CHAR(60) NOT NULL,
						firstname VARCHAR(255) NOT NULL,
						lastname VARCHAR(255) NOT NULL,
						PRIMARY KEY (id)
					 ) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
		{
			Version:     1.1,
			Description: "Creating table books",
			Script: `CREATE TABLE books (
						id INT 		auto_increment, 
						name 		VARCHAR(255) NOT NULL,
						description VARCHAR(255) NOT NULL,
						body MEDIUMTEXT NOT NULL,
						PRIMARY KEY (id)
					 ) ENGINE=InnoDB CHARACTER SET=utf8;`,
		},
	}
)

func Run(db *sql.DB) {
	log.Info("running migrations")
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations, nil)
	err := d.Migrate()
	if err != nil {
		log.Panic(err)
	}
}
