package TestData

import "github.com/fitan/genapi/pkg/gen_apiV2/TestData/nest"

type User struct {
	ID    int
	Age   int
	Name  string
	Nest  nest.Nest
	Fater nest.Fater
}
