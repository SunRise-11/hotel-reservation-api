// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/countries": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "findAll Countries",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "X-TenantID",
                        "name": "X-TenantID",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Country"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "create new Country",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "X-TenantID",
                        "name": "X-TenantID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Country",
                        "name": "country",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    }
                }
            }
        },
        "/countries/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "find Country by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "X-TenantID",
                        "name": "X-TenantID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "update Country",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "X-TenantID",
                        "name": "X-TenantID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Country",
                        "name": "Country",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    },
                    {
                        "description": "Country",
                        "name": "country",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Country"
                        }
                    }
                }
            }
        },
        "/countries/{id}/provinces": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Country"
                ],
                "summary": "find Provinces by country ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "X-TenantID",
                        "name": "X-TenantID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "Id",
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
                                "$ref": "#/definitions/models.Province"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.City": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "$ref": "#/definitions/models.Province"
                },
                "province_id": {
                    "type": "integer"
                },
                "tenant_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "models.Country": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "provinces": {
                    "description": "swagger:ignore",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Province"
                    }
                },
                "tenant_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "models.Province": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "cities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.City"
                    }
                },
                "country": {
                    "$ref": "#/definitions/models.Country"
                },
                "country_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "tenant_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
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
	Title:            "Reservation API",
	Description:      "Swagger documentation for reservation API .",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
