package TestData

import (
	"github.com/fitan/genapi/pkg/gin_api/TestData/nest"
	"time"
)

type UserResult struct {
	Code AliaseInt                 `json:"code"`
	Data map[string]nest.Nest      `json:"data"`
	M    map[nest.Fater]nest.Fater `json:"m"`
	N    nest.Nest                 `json:"n"`
	Err  interface{}               `json:"err"`
	User User
	//ATime time.Time
}

type AliaseInt int

type User struct {
	ID         int `json:"id"`
	Age        int `json:"age"`
	Name       string
	Nest       []nest.Nest          `json:"nest"`
	M          map[string]nest.Nest `json:"m"`
	Fater      nest.Fater           `json:"fater"`
	Time       time.Time            `json:"time"`
	UserResult *UserResult
}

type Test1T []nest.Nest
