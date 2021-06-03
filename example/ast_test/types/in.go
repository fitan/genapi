package types

import (
	models2 "ast_test/models"
)

type alia string

type alia2 struct {
	F1 string
}

type In struct {
	St struct {
		// this is h_1
		H1 string `json:"h_1"`
		// 这是h_2
		H2 string `json:"h_2"`
	}
	M  map[string]string
	M2 map[string]alia
	Mt map[string]alia2
	A  alia
	A2 alia2
	alia2
	models2.Out
	Out1 models2.Out
	// 你好啊
	Hello string `json:"hello"`
}