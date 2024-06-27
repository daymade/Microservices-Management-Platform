// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "daymade",
            "email": "daymadev89@gmail.com"
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
        "/services": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "获取服务列表，支持分页、排序和搜索",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "services"
                ],
                "summary": "列出服务",
                "parameters": [
                    {
                        "type": "string",
                        "description": "搜索关键词",
                        "name": "query",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序字段",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序方向 (asc 或 desc)",
                        "name": "sort_direction",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页条数",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/viewmodel.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/viewmodel.ServiceListViewModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/services/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "通过 ID 获取单个服务的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "services"
                ],
                "summary": "获取单个服务",
                "parameters": [
                    {
                        "type": "string",
                        "description": "服务 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/viewmodel.ServiceDetailViewModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "返回当前登录用户的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "获取当前登录用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/viewmodel.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "viewmodel.PaginatedResponse": {
            "description": "分页响应模型，包含分页数据和分页信息",
            "type": "object",
            "properties": {
                "data": {},
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "page_size": {
                    "type": "integer",
                    "example": 20
                },
                "total_count": {
                    "type": "integer",
                    "example": 100
                },
                "total_pages": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "viewmodel.ServiceDetailViewModel": {
            "description": "服务详情模型，包含服务的详细信息和版本列表",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "This is a sample service"
                },
                "id": {
                    "type": "string",
                    "example": "srv-123"
                },
                "name": {
                    "type": "string",
                    "example": "My Service"
                },
                "owner_id": {
                    "type": "string",
                    "example": "usr-456"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-02T00:00:00Z"
                },
                "versions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/viewmodel.VersionViewModel"
                    }
                }
            }
        },
        "viewmodel.ServiceListViewModel": {
            "description": "服务列表项模型，包含服务的基本信息",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "This is a sample service"
                },
                "id": {
                    "type": "string",
                    "example": "srv-123"
                },
                "name": {
                    "type": "string",
                    "example": "My Service"
                },
                "owner_id": {
                    "type": "string",
                    "example": "usr-456"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-02T00:00:00Z"
                },
                "version_count": {
                    "type": "integer",
                    "example": 3
                }
            }
        },
        "viewmodel.User": {
            "description": "用户信息模型，包含用户的基本信息",
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "https://example.com/avatar.jpg"
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "username": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        },
        "viewmodel.VersionViewModel": {
            "description": "版本信息模型，包含版本的基本信息",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00Z"
                },
                "description": {
                    "type": "string",
                    "example": "Initial release"
                },
                "number": {
                    "type": "string",
                    "example": "v1.0.0"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Catalog Service Management",
	Description:      "This is a platform to manage services.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
