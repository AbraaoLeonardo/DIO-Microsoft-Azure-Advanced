package db

import (
	"database/sql"
	"fmt"
	"server/utils"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	env := utils.GetEnv()
	connectString := fmt.Sprintf("postgresql://%s@%s@%s:%s/%s", env.Db_User, env.Db_Pass, env.Db_Host, env.Db_Port, env.Db_Table)

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) CloseDatabase() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
