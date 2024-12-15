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
        "/classes": {
            "get": {
                "description": "Retrieves a list of classes with pagination.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Classes"
                ],
                "summary": "List classes with pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Class"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new class with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Classes"
                ],
                "summary": "Create a new class",
                "parameters": [
                    {
                        "description": "Class details",
                        "name": "class",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ClassView"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Class"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/classes/{id}": {
            "get": {
                "description": "Retrieves a class by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Classes"
                ],
                "summary": "Get a class by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Class ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ClassView"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Updates the details of an existing class.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Classes"
                ],
                "summary": "Update an existing class",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Class ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Class details",
                        "name": "class",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ClassView"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Class"
                        }
                    },
                    "400": {
                        "description": "Invalid input or class is used in services",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a class by its ID. If the class is used in any services, it returns an error.",
                "tags": [
                    "Classes"
                ],
                "summary": "Delete a class",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Class ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Class deleted successfully"
                    },
                    "400": {
                        "description": "Class is used in services",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/parameters": {
            "get": {
                "description": "Retrieves a list of parameters with pagination.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parameters"
                ],
                "summary": "List parameters with pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Parameter"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new parameter with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parameters"
                ],
                "summary": "Create a new parameter",
                "parameters": [
                    {
                        "description": "Parameter details",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParameterView"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Parameter"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/parameters/{id}": {
            "get": {
                "description": "Retrieves a parameter by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parameters"
                ],
                "summary": "Get a parameter by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Parameter ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ParameterView"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an existing parameter with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parameters"
                ],
                "summary": "Update an existing parameter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Parameter ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Parameter details",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParameterView"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Parameter"
                        }
                    },
                    "400": {
                        "description": "Invalid input or parameter is used in services",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a parameter by its ID. If the parameter is used in any services, it returns an error.",
                "tags": [
                    "Parameters"
                ],
                "summary": "Delete a parameter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Parameter ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Parameter deleted successfully"
                    },
                    "400": {
                        "description": "Parameter is used in services",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/report": {
            "get": {
                "description": "Generates a fiscal report in PDF format and returns it as a downloadable file.",
                "produces": [
                    "application/pdf"
                ],
                "tags": [
                    "Reports"
                ],
                "summary": "Build fiscal report",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/services": {
            "get": {
                "description": "Fetches a list of services with pagination.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "List all services",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Service"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new service with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Create a new service",
                "parameters": [
                    {
                        "description": "Service details",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.NewService"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Service"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/services/{id}": {
            "get": {
                "description": "Fetches the details of a service by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Get a service by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Service"
                        }
                    },
                    "404": {
                        "description": "Service not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/services/{id}/approve": {
            "post": {
                "description": "Approves a service by its ID. If a class ID is provided in the request body, it assigns the class to the service before approval.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Approve a service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Class ID",
                        "name": "class",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.assignClassRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Service"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Service or class not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/services/{id}/proposed_classes": {
            "get": {
                "description": "Fetches a list of proposed classes for a service based on similar parameters.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "List proposed classes for a service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
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
                                "$ref": "#/definitions/models.Class"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid service ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Service not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.NewService": {
            "type": "object",
            "required": [
                "parameters",
                "title"
            ],
            "properties": {
                "parameters": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handlers.assignClassRequest": {
            "type": "object",
            "properties": {
                "class_id": {
                    "type": "integer"
                }
            }
        },
        "models.Class": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ClassView": {
            "type": "object",
            "properties": {
                "allowed_parameters": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "mob_inet",
                        "fix_ctv",
                        "voice_fix"
                    ]
                },
                "id": {
                    "type": "integer",
                    "example": 3042
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Parameter": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ParameterView": {
            "type": "object",
            "properties": {
                "allowed_classes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        1033,
                        3023
                    ]
                },
                "contradiction_parameters": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "mob_inet",
                        "fix_ctv",
                        "voice_fix"
                    ]
                },
                "id": {
                    "type": "string",
                    "example": "fix_ctv"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Service": {
            "type": "object",
            "properties": {
                "approved_at": {
                    "type": "string"
                },
                "class": {
                    "$ref": "#/definitions/models.Class"
                },
                "class_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Parameter"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "194.135.25.202:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "MyApp API",
	Description:      "This is a backend server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
