package utils

import "os"


type Postgresql struct {
	Db_Port		string
	Db_User		string
	Db_Host		string
	Db_Table	string
	Db_Pass		string
}

func GetEnv() Postgresql{
	return Postgresql{
		Db_Host: "postgres",
		Db_User: os.Getenv("POSTGRES_USER"),
		Db_Port: os.Getenv("POSTGRES_PORT"),
		Db_Pass: os.Getenv("POSTGRES_PASSWORD"),
		Db_Table: os.Getenv("tarefas"),
	}
}
