package models

import "ast_test/newmodelsv2"

type Out struct {
	Data struct{
		Name string `json:"name"`
		Age string `json:"age"`
		Sub newmodelsv2.SubObj `json:"sub"`
	}
}
