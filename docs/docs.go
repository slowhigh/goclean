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
        "/v1/memo": {
            "get": {
                "description": "list memos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memo"
                ],
                "summary": "List memo",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"2021-02-18T21:54:42.123Z\"",
                        "description": "createAt search by start",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"2021-02-18T21:54:42.123Z\"",
                        "description": "createAt search by end",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "content search by keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "content search by keyword",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.FindAllMemoRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "create memo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memo"
                ],
                "summary": "Create memo",
                "parameters": [
                    {
                        "description": "writer and content for the new memo",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.CreateMemoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.CreateMemoRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/v1/memo/{id}": {
            "get": {
                "description": "get memo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memo"
                ],
                "summary": "Get memo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID search by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.FindOneMemoRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "update memo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memo"
                ],
                "summary": "Update memo",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "memo update by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "writer and content for the new memo",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.UpdateMemoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.UpdateMemoRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "delete memo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memo"
                ],
                "summary": "Delete memo",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "memo delete by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.DeleteMemoRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.CreateMemoReq": {
            "type": "object",
            "required": [
                "content",
                "writer"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 1
                },
                "writer": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 1
                }
            }
        },
        "github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.CreateMemoRes": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.DeleteMemoRes": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "update_at": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.FindAllMemoRes": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "update_at": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.FindOneMemoRes": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "update_at": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.UpdateMemoReq": {
            "type": "object",
            "required": [
                "content",
                "writer"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 1
                },
                "writer": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 1
                }
            }
        },
        "github_com_slowhigh_goclean_internal_adapter_controller_rest_dto_memo_dto.UpdateMemoRes": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "update_at": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "goclean",
	Description:      "goclean",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}