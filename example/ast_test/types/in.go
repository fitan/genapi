package types

import (
	models "ast_test/models"
)
type alia string

type alia2 struct {
	F1 string
}

type In struct {
	St struct{
		H1 string
		H2 string
	}
	M map[string]string
	M2 map[string]alia
	Mt map[string]alia2
	A alia
	A2 alia2
	models.Out
	Out1 models.Out
	Hello string `json:"hello"`
}