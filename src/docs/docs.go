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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/about": {
            "get": {
                "description": "Returns the supported API versions and the internal build nr",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Lists general information about the API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/datastructures.About"
                        }
                    }
                }
            }
        },
        "/v1/groups/{number}": {
            "get": {
                "description": "List all Signal Groups.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "List all Signal Groups.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/datastructures.GroupEntry"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Signal Group with the specified members.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Create a new Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/datastructures.CreateGroup"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            }
        },
        "/v1/groups/{number}/{groupid}": {
            "delete": {
                "description": "Delete a Signal Group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Delete a Signal Group.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group Id",
                        "name": "groupid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            }
        },
        "/v1/qrcodelink": {
            "get": {
                "description": "test",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Link device and generate QR code.",
                "responses": {
                    "200": {
                        "description": "Image",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/receive/{number}": {
            "get": {
                "description": "Receives Signal Messages from the Signal Network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Receive Signal Messages.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            }
        },
        "/v1/register/{number}": {
            "post": {
                "description": "Register a phone number with the signal network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Register a phone number.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            }
        },
        "/v1/register/{number}/verify/{token}": {
            "post": {
                "description": "Verify a registered phone number with the signal network.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Verify a registered phone number.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registered Phone Number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Additional Settings",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datastructures.VerifyNumberSettings"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Verification Code",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            }
        },
        "/v1/send": {
            "post": {
                "description": "Send a signal message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Send a signal message.",
                "deprecated": true,
                "parameters": [
                    {
                        "description": "Input Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datastructures.SendMessageV1"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            }
        },
        "/v2/send": {
            "post": {
                "description": "Send a signal message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messages"
                ],
                "summary": "Send a signal message.",
                "parameters": [
                    {
                        "description": "Input Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/datastructures.SendMessageV2"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/datastructures.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "datastructures.About": {
            "type": "object",
            "properties": {
                "build": {
                    "type": "integer"
                },
                "versions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "datastructures.CreateGroup": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "datastructures.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "datastructures.GroupEntry": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "blocked": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "internal_id": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "datastructures.SendMessageV1": {
            "type": "object",
            "properties": {
                "base64_attachment": {
                    "type": "string"
                },
                "is_group": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "datastructures.SendMessageV2": {
            "type": "object",
            "properties": {
                "base64_attachments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "datastructures.VerifyNumberSettings": {
            "type": "object",
            "properties": {
                "pin": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "List general information.",
            "name": "General"
        },
        {
            "description": "Register and link Devices.",
            "name": "Devices"
        },
        {
            "description": "Create, List and Delete Signal Groups.",
            "name": "Groups"
        },
        {
            "description": "Send and Receive Signal Messages.",
            "name": "Messages"
        }
    ]
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
	Host:        "127.0.0.1:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Signal Cli REST API",
	Description: "This is the Signal Cli REST API documentation.",
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
