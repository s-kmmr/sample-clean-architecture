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
        "termsOfService": "none",
        "contact": {},
        "license": {
            "name": "kmmr"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/members": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "member"
                ],
                "summary": "会員情報一覧をリストで返す",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.MemberResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
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
                    "member"
                ],
                "summary": "会員情報の登録",
                "parameters": [
                    {
                        "description": "Add New Member",
                        "name": "member",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MemberRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.MemberRequest": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.MemberResponse": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "response.Success": {
            "type": "object",
            "properties": {
                "message": {
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "sample-clean-architecture API",
	Description: "sample-clean-architecture OpenAPI",
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
