// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/car": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "create one car",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genrest.CarNodeNotID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.CarNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/car/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "get one car",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.CarNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "update one car",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genrest.CarNodeNotID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.CarNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "delete one car",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.CarNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/cars": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "get list car",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.GetCarListData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "update list car",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/genrest.CarNode"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.CarNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "create list car",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/genrest.CarNode"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/genrest.CarNodeNotID"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car"
                ],
                "summary": "delete list car",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.CarNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create one user",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genrest.UserNodeNotID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.UserNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get one user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.UserNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update one user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genrest.UserNodeNotID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.UserNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete one user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.UserNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/{id}/cars": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car by User"
                ],
                "summary": "get list car",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/genrest.CarNode"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car by User"
                ],
                "summary": "create list car",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/genrest.CarNodeNotID"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/genrest.CarNodeNotID"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Car by User"
                ],
                "summary": "delete one car",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.CarNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get list user",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "eq_age1",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.GetUserListData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update list user",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/genrest.UserNode"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.UserNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create list user",
                "parameters": [
                    {
                        "description": " ",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/genrest.UserNode"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/genrest.UserNodeNotID"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete list user",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "name": "ids",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/genrest.RestReturn"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/genrest.UserNode"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Car": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "model": {
                    "description": "Model holds the value of the \"model\" field.",
                    "type": "string"
                },
                "registered_at": {
                    "description": "RegisteredAt holds the value of the \"registered_at\" field.",
                    "type": "string"
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "age1": {
                    "description": "Age1 holds the value of the \"age1\" field.",
                    "type": "integer"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the UserQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.UserEdges"
                },
                "en": {
                    "description": "En holds the value of the \"en\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.UserEdges": {
            "type": "object",
            "properties": {
                "cars": {
                    "description": "Cars holds the value of the cars edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Car"
                    }
                }
            }
        },
        "genrest.CarNode": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "model": {
                    "type": "string"
                },
                "registered_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "genrest.CarNodeNotID": {
            "type": "object",
            "properties": {
                "model": {
                    "type": "string"
                },
                "registered_at": {
                    "type": "string",
                    "format": "date-time"
                }
            }
        },
        "genrest.GetCarListData": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Car"
                    }
                }
            }
        },
        "genrest.GetUserListData": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.User"
                    }
                }
            }
        },
        "genrest.RestReturn": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "object"
                }
            }
        },
        "genrest.UserNode": {
            "type": "object",
            "properties": {
                "age1": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "genrest.UserNodeNotID": {
            "type": "object",
            "properties": {
                "age1": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
