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
        "license": {
            "name": "MIT",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chapter/{manga}/{chapter}": {
            "get": {
                "description": "Scrape page urls of the chapter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chapter"
                ],
                "summary": "Get schapter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested manga",
                        "name": "manga",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Requested chapter",
                        "name": "chapter",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/history": {
            "get": {
                "description": "Delete record from the history",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "history"
                ],
                "summary": "Remove from history",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/library": {
            "get": {
                "description": "Obtain mangas in the library",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Get from library",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create new manga in the library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Add to library",
                "parameters": [
                    {
                        "description": "Requested add",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AddLibraryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Delete manga from the library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Remove from library",
                "parameters": [
                    {
                        "description": "Requested delete",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RemoveLibraryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/manga/{manga}": {
            "get": {
                "description": "Obtain details about manga",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manga"
                ],
                "summary": "Get manga",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested manga",
                        "name": "manga",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/progress": {
            "get": {
                "description": "Obtain user progress on chapter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "progress"
                ],
                "summary": "Get progress",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested progress",
                        "name": "manga",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Requested chapter",
                        "name": "chapter",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Renew current manga reading status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "progress"
                ],
                "summary": "Update progress",
                "parameters": [
                    {
                        "description": "Requested progress",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateProgressRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "Search for manga by title",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manga"
                ],
                "summary": "Search manga",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested title",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Register user via Telegram data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Requested user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "types.AddLibraryRequest": {
            "type": "object",
            "properties": {
                "cover_image": {
                    "type": "string"
                },
                "manga": {
                    "type": "string"
                }
            }
        },
        "types.CreateUserRequest": {
            "type": "object",
            "required": [
                "first_name",
                "id"
            ],
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "types.RemoveLibraryRequest": {
            "type": "object",
            "properties": {
                "manga": {
                    "type": "string"
                }
            }
        },
        "types.UpdateProgressRequest": {
            "type": "object",
            "properties": {
                "chapter": {
                    "type": "string"
                },
                "manga": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
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
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Yomu API",
	Description:      "Yomu is a free manga reader Telegram mini app.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
