// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/healthcheck": {
            "get": {
                "description": "Get server status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Show server status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.HealthCheckResponse"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Authenticates a user using email and password, and generates a new JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Logs in a user",
                "parameters": [
                    {
                        "description": "The email and password of the user",
                        "name": "InputBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns a struct with the JWT and its expiration timestamp",
                        "schema": {
                            "$ref": "#/definitions/schemas.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Returns an error message if the request body cannot be parsed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Returns an error message if the provided password does not match the hash stored in the database",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Returns an error message if there is a server-side issue",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "description": "get many users based on pagination and sorting parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get many users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sorting (format: field[direction])",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Register a new user with email, screen_name, and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "description": "Retrieve a user from the database by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "406": {
                        "description": "Multiple users found",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "companyId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "rank": {
                    "type": "integer"
                },
                "screenName": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "schemas.CreateUserRequest": {
            "description": "This request includes the necessary details to create a new user.",
            "type": "object",
            "required": [
                "email",
                "password",
                "screen_name"
            ],
            "properties": {
                "company_id": {
                    "description": "The ID of the company to associate with the new user. This field is optional and must be greater than 0 if provided.\n@Min 0\n@Example 1",
                    "type": "integer"
                },
                "email": {
                    "description": "The email address of the new user. This field is required and must be a valid email address.\n@Format email\n@Example user@example.com",
                    "type": "string"
                },
                "password": {
                    "description": "The password for the new user. This field is required and must contain at least 7 characters.\n@MinLength 7\n@Example 1234567",
                    "type": "string",
                    "minLength": 7
                },
                "rank": {
                    "description": "The rank of the new user. This field is optional and must be greater than 0 if provided.\n@Min 0\n@Example 1",
                    "type": "integer"
                },
                "screen_name": {
                    "description": "The desired username for the new user. This field is required and must contain at least 4 characters.\n@MinLength 4\n@Example exampleUser",
                    "type": "string",
                    "minLength": 4
                }
            }
        },
        "schemas.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "schemas.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "schemas.LoginInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 7
                }
            }
        },
        "schemas.TokenResponse": {
            "type": "object",
            "properties": {
                "access_expires": {
                    "type": "integer"
                },
                "access_token": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v.1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Auth Service API",
	Description:      "API server for auth stuff",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
