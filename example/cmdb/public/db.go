package public

import (
	"cmdb/ent"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *ent.Client

func Init() {
	DB = NewDB()
}

func NewDB() *ent.Client {
	db, err := ent.Open("mysql", GetConf().Mysql.Addr)
	if err != nil {
		XLog.Error().Err(err)
	}
	if GetConf().Mysql.Debug {
		db.Debug()
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		XLog.Error().Err(fmt.Errorf("failed creating schema resources: %v", err))
	}
	return db
}

func GetDB() *ent.Client {
	if DB == nil {
		DB = NewDB()
	}

	return DB
}
