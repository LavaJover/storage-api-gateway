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
        "/auth/login": {
            "post": {
                "description": "Login user using email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.authOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Wrong credentials",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "SSO service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Refresh access-JWT using refresh-JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "refresh"
                ],
                "summary": "Refresh access-JWT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Refresh-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.authOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Invalid token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "SSO service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/reg": {
            "post": {
                "description": "Register new user using email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "signup"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.registerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.authOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method is not supported",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Email is already taken",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "SSO service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/valid": {
            "post": {
                "description": "Validate JWT passed in HTTP POST request header by Bearer scheme",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "valid"
                ],
                "summary": "Validate JWT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.validateJWTOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Invalid token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "SSO service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/boxes": {
            "get": {
                "description": "Get all boxes related to given cell_id checking user permissions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boxes"
                ],
                "summary": "Get boxes by cell_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Cell ID",
                        "name": "cell_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.getBoxesOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "You don't have enough permissions",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method is not supported",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Storage service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new named box connected to cell",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boxes"
                ],
                "summary": "Create new box",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.createBoxOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method is not supported",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Storage service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cells": {
            "get": {
                "description": "Get all cells by given storage_id with permission checking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cells"
                ],
                "summary": "Get cells by storage_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Storage ID",
                        "name": "storage_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.getCellsOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "You don't have enough permissions",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method is not supported",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Storage service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new named cell connected to storage to store boxes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cells"
                ],
                "summary": "Create new cell",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.createCellOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method is not supported",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Storage service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/storages": {
            "get": {
                "description": "Get all storage instances related to the given user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storages"
                ],
                "summary": "Get storages by user_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.getStoragesOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "You dont't have enough permissions",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method is not supported",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Storage service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new named storage to store cells",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storages"
                ],
                "summary": "Create new storage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Access-JWT",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.createStorageOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method is not supported",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Storage service failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.authOkResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "main.createBoxOkResponse": {
            "type": "object",
            "properties": {
                "cell_id": {
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
        "main.createCellOkResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "storage_id": {
                    "type": "integer"
                }
            }
        },
        "main.createStorageOkResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "main.getBoxesOkResponse": {
            "type": "object",
            "properties": {
                "boxes": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "cell_id": {
                                "type": "integer"
                            },
                            "id": {
                                "type": "integer"
                            },
                            "name": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "main.getCellsOkResponse": {
            "type": "object",
            "properties": {
                "cells": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "id": {
                                "type": "integer"
                            },
                            "name": {
                                "type": "string"
                            },
                            "storage_id": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "main.getStoragesOkResponse": {
            "type": "object",
            "properties": {
                "storages": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "id": {
                                "type": "integer"
                            },
                            "name": {
                                "type": "string"
                            },
                            "user_id": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "main.loginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "main.registerRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "main.validateJWTOkResponse": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Storage-Master API",
	Description:      "API for storage-api-gateway",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
