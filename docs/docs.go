// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-10 16:32:13.973348 -0700 PDT m=+0.053975305

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Binh Nguyen",
            "url": "http://www.swagger.io/support",
            "email": "ntbinh106@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/kubectl": {
            "get": {
                "description": "get kubernetes objects with kubectl",
                "produces": [
                    "application/json"
                ],
                "summary": "execute kubectl command",
                "parameters": [
                    {
                        "type": "string",
                        "description": "use by kubectl",
                        "name": "cmd",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kubechat.CmdOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "kubechat.CmdOutput": {
            "type": "object",
            "properties": {
                "cmd": {
                    "type": "string"
                },
                "code": {
                    "type": "integer"
                },
                "responseTime": {
                    "type": "string"
                },
                "stderr": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "stdout": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
	Version:     "0.1.0",
	Host:        "localhost:8080",
	BasePath:    "/v1",
	Schemes:     []string{"http", "https"},
	Title:       "Swagger Kubechat API",
	Description: "This is a API wraper of kubectl for chatbot",
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
