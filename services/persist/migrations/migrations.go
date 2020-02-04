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
