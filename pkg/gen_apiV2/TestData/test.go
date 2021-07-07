package TestData

import (
	"github.com/fitan/genapi/pkg/gen_apiV2/TestData/nest"
	"time"
)

type User struct {
	ID    int `json:"id"`
	Age   int `json:"age"`
	Name  string
	Nest  nest.Nest `json:"nest"`
	Fater nest.Fater `json:"fater"`
	Time time.Time `json:"time"`
}
