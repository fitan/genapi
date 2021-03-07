package public

import (
	"ent_samp/ent"
	"log"
)

var DB *ent.Client

func Init()  {
	DB = NewDB()
}

func NewDB() *ent.Client {
	db, err := ent.Open("mysql", "root:123456@tcp(localhost:3306)/ent?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
	}
	return db
}

func GetDB() *ent.Client {
	if DB == nil {
		DB = NewDB()
		return DB
	}

	return DB
}