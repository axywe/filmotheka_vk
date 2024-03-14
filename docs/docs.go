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
        "/actors": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actors"
                ],
                "summary": "Get list of actors",
                "responses": {
                    "200": {
                        "description": "List of actors",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/actor.Actor"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actors"
                ],
                "summary": "Update an actor",
                "parameters": [
                    {
                        "description": "Actor with updated information",
                        "name": "actor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/actor.Actor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Actor updated",
                        "schema": {
                            "$ref": "#/definitions/actor.Actor"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actors"
                ],
                "summary": "Create a new actor",
                "parameters": [
                    {
                        "description": "Actor to create",
                        "name": "actor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/actor.Actor"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Actor created",
                        "schema": {
                            "$ref": "#/definitions/actor.Actor"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Actors"
                ],
                "summary": "Delete an actor",
                "responses": {
                    "200": {
                        "description": "Actor deleted"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/auth": {
            "post": {
                "description": "Processes POST user authentication requests and generates JWT tokens.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authentication Processing",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful authentication",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "User not found or invalid credentials",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Only the POST method is allowed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Get list of movies",
                "responses": {
                    "200": {
                        "description": "List of movies",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/movie.Movie"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Update a movie",
                "parameters": [
                    {
                        "description": "Movie with updated information",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/movie.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Movie updated",
                        "schema": {
                            "$ref": "#/definitions/movie.Movie"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Create a new movie",
                "parameters": [
                    {
                        "description": "Movie to create",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/movie.Movie"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Movie created",
                        "schema": {
                            "$ref": "#/definitions/movie.Movie"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Delete a movie",
                "responses": {
                    "200": {
                        "description": "Movie deleted"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "actor.Actor": {
            "type": "object",
            "properties": {
                "birthdate": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "movies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/actor.MovieBrief"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "actor.MovieBrief": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "auth.Credentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "movie.Movie": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "releaseDate": {
                    "type": "string"
                },
                "title": {
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
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Filmotheka API",
	Description:      "This is a server for Filmotheka application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}