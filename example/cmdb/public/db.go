package public

import (
	"cmdb/ent"
	"context"
	"log"
)

var DB *ent.Client

func Init() {
	DB = NewDB()
}

func NewDB() *ent.Client {
	db, err := ent.Open("mysql", "root:123456@tcp(10.143.131.148:3306)/cmdb?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
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
