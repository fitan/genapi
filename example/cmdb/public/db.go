package public

import (
	"cmdb/ent"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var db *ent.Client
var dbLock sync.Mutex

func NewDB() (*ent.Client, error){
	db, err := ent.Open("mysql", GetConf().Mysql.Addr)
	if err != nil {
		return nil, err
	}
	if GetConf().Mysql.Debug {
		db.Debug()
	}
	//if err := db.Schema.Create(context.Background()); err != nil {
	//	XLog.Error().Msgf("failed creating schema resources: %v", err)
	//}
	return db, nil
}

func GetDB() *ent.Client {
	if db == nil {
		dbLock.Lock()
		defer dbLock.Unlock()
		d, err := NewDB()
		if err != nil {
			GetXLog().Error().Err(err).Msg("")
		}
		db = d
	}
	return db
}
