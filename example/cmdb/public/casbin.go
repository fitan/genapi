package public

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"sync"
)

var enforcer *casbin.Enforcer
var casbinLock sync.Mutex

type Header struct {
	// 这是Header 的姓名
	Name string `header:"name"`
	// 这是age 的姓名
	Age int `header:"age"`
	HeaderSub HeaderSub

}

type HeaderSub struct {
	// 这是age sub 的姓名
	AgeSub int `header:"ageSub"`
}

func newCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapter("mysql", GetConf().Mysql.Addr, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("conf/rbac_model.conf", a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}

func GetCasbin() *casbin.Enforcer {
	if enforcer == nil {
		casbinLock.Lock()
		defer casbinLock.Unlock()
		e, err := newCasbin()
		if err != nil {
			GetXLog().Error().Err(err).Msg("")
		}
		enforcer = e
	}
	return enforcer
}

func CheckKeysCasbin(c *gin.Context, casbinMark string,key interface{}) (bool, error) {
	userName,has := c.Get("user_name")
	if !has {
		return has, errors.New("not found key: userName")
	}
	v := []interface{}{userName, key,casbinMark}
	has, err := GetCasbin().Enforce(v...)
	if err != nil {
		return false, err
	}

	if !has {
		return false, errors.New("no permission")
	}
	return true, nil
}

func CheckListKeysCasbin(c *gin.Context, casbinMark string, keys [][]interface{}) (bool, error)  {
	userName, has := c.Get("user_name")
	if !has {
		return has, errors.New("not found key: userName")
	}
	vs := make([][]interface{},0,0)
	for _, key := range keys {
		v := []interface{}{userName}
		v = append(v, key...)
		v = append(v, casbinMark)
		vs = append(vs, v)
	}
	enforce, err := GetCasbin().BatchEnforce(vs)
	if err != nil {
		return false, err
	}
	errs := make([]string,0,0)
	for index, has := range enforce {
		if !has {
			errs = append(errs,fmt.Sprintf("id: %s permission denied", vs[index][1]))
		}
	}
	if len(errs) != 0 {
		return true, errors.New(strings.Join(errs, "\n"))
	}
	return true, nil
}

type GetCasbinKeyser interface {
	GetCasbinKeys() []interface{}
}
