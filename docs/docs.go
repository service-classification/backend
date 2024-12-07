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
        "/groups": {
            "get": {
                "description": "Retrieves a list of groups with pagination.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "List groups with pagination",
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
                                "$ref": "#/definitions/models.Group"
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
                "description": "Creates a new group with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Create a new group",
                "parameters": [
                    {
                        "description": "Group details",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.NewGroup"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
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
        "/groups/{id}": {
            "put": {
                "description": "Updates the details of an existing group.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Groups"
                ],
                "summary": "Update an existing group",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Group details",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.NewGroup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Group"
                        }
                    },
                    "400": {
                        "description": "Invalid input or group is used in services",
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
                "description": "Deletes a group by its ID. If the group is used in any services, it returns an error.",
                "tags": [
                    "Groups"
                ],
                "summary": "Delete a group",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Group deleted successfully"
                    },
                    "400": {
                        "description": "Group is used in services",
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
                            "$ref": "#/definitions/services.NewParameter"
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
                            "$ref": "#/definitions/services.NewParameter"
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
        "/services/{id}/approve": {
            "post": {
                "description": "Approves a service by its ID. If a group ID is provided in the request body, it assigns the group to the service before approval.",
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
                        "description": "Group ID",
                        "name": "group",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.assignGroupRequest"
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
                        "description": "Service or group not found",
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
        "/services/{id}/proposed_groups": {
            "get": {
                "description": "Fetches a list of proposed groups for a service based on similar parameters.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "List proposed groups for a service",
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
                                "$ref": "#/definitions/models.Group"
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
        "handlers.assignGroupRequest": {
            "type": "object",
            "properties": {
                "group_id": {
                    "type": "integer"
                }
            }
        },
        "models.Group": {
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
        "models.Service": {
            "type": "object",
            "properties": {
                "approved_at": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "group": {
                    "$ref": "#/definitions/models.Group"
                },
                "group_id": {
                    "type": "integer"
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
        },
        "services.NewGroup": {
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
        "services.NewParameter": {
            "type": "object",
            "properties": {
                "allowed_classes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "1",
                        "1033",
                        "3023"
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
