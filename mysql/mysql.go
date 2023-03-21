package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDB() {
	
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/operant")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
 		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://mysql/migrations",
		"mysql",
		driver,
	)
	
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
