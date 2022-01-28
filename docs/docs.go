// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "ha 1.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/autoCodeExample/createAutoCodeExample": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRegister"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户模型",
                        "name": "userReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建用户",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/user.UserRes"
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
        "user.UserReq": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 10
                },
                "birthday": {
                    "type": "string",
                    "example": "0001-01-01 00:00:00"
                },
                "company": {
                    "type": "string",
                    "example": "小明"
                },
                "department": {
                    "type": "string",
                    "example": "小明"
                },
                "info": {
                    "type": "string",
                    "example": "小明"
                },
                "mail": {
                    "type": "string",
                    "example": "小明"
                },
                "name": {
                    "type": "string",
                    "example": "小明"
                },
                "password": {
                    "type": "string",
                    "example": "小明"
                },
                "qqnumber": {
                    "type": "string",
                    "example": "小明"
                },
                "realname": {
                    "type": "string",
                    "example": "小明"
                },
                "sex": {
                    "type": "string",
                    "example": "小明"
                },
                "vocation": {
                    "type": "string",
                    "example": "小明"
                },
                "wechat": {
                    "type": "string",
                    "example": "小明"
                }
            }
        },
        "user.UserRes": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "birthday": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "createDate": {
                    "type": "string"
                },
                "department": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "qqNumber": {
                    "type": "string"
                },
                "realName": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "vocation": {
                    "type": "string"
                },
                "wechat": {
                    "type": "string"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
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
	Host:        "47.107.108.127:80",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "JieTong API",
	Description: "除了用户登录和注册以及头像获取三个接口\n其他的都需要用户携带TOKEN进行用户验证，否则无法访问接口\n\nTOKEN 格式 ： KEY：Authorization VALUE： \"JWT \" + 登录时返回的对应TOKEN   （放在请求的header中）\n",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
