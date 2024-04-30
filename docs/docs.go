// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/invention": {
            "get": {
                "description": "Get all inventions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/v1"
                ],
                "summary": "Get all inventions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Invention"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new invention with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/v1"
                ],
                "summary": "Add a new invention",
                "parameters": [
                    {
                        "description": "Invention entity",
                        "name": "invention",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Invention"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data is valid",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/invention/{id}": {
            "get": {
                "description": "Get an invention by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api/v1"
                ],
                "summary": "Get an invention by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Invention ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Invention"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Invention": {
            "type": "object",
            "required": [
                "inventor"
            ],
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2024-04-30T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "maxLength": 300,
                    "minLength": 3,
                    "example": "This is an amazing invention that will revolutionize the world."
                },
                "id": {
                    "type": "string",
                    "example": "123"
                },
                "inventor": {
                    "$ref": "#/definitions/entity.Person"
                },
                "title": {
                    "type": "string",
                    "minLength": 3,
                    "example": "Amazing Invention"
                }
            }
        },
        "entity.Person": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 18,
                    "example": 30
                },
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "firstname": {
                    "type": "string",
                    "example": "John"
                },
                "lastname": {
                    "type": "string",
                    "example": "Doe"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.7.3",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Mini-Go-Server",
	Description:      "API Server for my github project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}