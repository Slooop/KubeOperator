// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-19 17:20:34.326379 +0800 CST m=+2.794629819

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
        "termsOfService": "http://kubeoperator.io",
        "contact": {
            "name": "Fit2cloud Support",
            "url": "https://www.fit2cloud.com",
            "email": "support@fit2cloud.com"
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
        "/clusters/": {
            "get": {
                "description": "List clusters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page num",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.ListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cluster",
                "parameters": [
                    {
                        "description": "cluster",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/serializer.CreateResponse"
                        }
                    }
                }
            }
        },
        "/clusters/batch/": {
            "post": {
                "description": "Batch Clusters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cluster",
                "parameters": [
                    {
                        "description": "Batch",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.BatchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.BatchResponse"
                        }
                    }
                }
            }
        },
        "/clusters/{cluster_name}": {
            "get": {
                "description": "Get Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.GetResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.DeleteResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a Cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cluster",
                "parameters": [
                    {
                        "description": "cluster",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UpdateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.BatchRequest": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.Cluster"
                    }
                },
                "operation": {
                    "type": "string"
                }
            }
        },
        "serializer.BatchResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.Cluster"
                    }
                }
            }
        },
        "serializer.Cluster": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "serializer.CreateRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "serializer.CreateResponse": {
            "type": "object",
            "properties": {
                "item": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.Cluster"
                }
            }
        },
        "serializer.DeleteResponse": {
            "type": "object"
        },
        "serializer.GetResponse": {
            "type": "object",
            "properties": {
                "item": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.Cluster"
                }
            }
        },
        "serializer.ListResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.Cluster"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "serializer.UpdateRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "serializer.UpdateResponse": {
            "type": "object",
            "properties": {
                "item": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.Cluster"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "KubeOperator Restful API",
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
