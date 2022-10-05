// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server that handle cute Gophers.",
    "title": "gophers-api",
    "version": "0.1.0"
  },
  "paths": {
    "/gopher": {
      "get": {
        "description": "Get a gopher by a given name",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "Gopher name",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A gopher",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Gopher"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "summary": "Add a new Gopher",
        "parameters": [
          {
            "description": "The Gopher to create.",
            "name": "gopher",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name",
                "path",
                "url"
              ],
              "properties": {
                "name": {
                  "type": "string"
                },
                "path": {
                  "type": "string"
                },
                "url": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Gopher"
            }
          }
        }
      }
    },
    "/gophers": {
      "get": {
        "description": "List Gophers",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "Return the Gophers list.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Gopher"
              }
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message.",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Gopher": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "example": "my-gopher"
        },
        "path": {
          "type": "string",
          "example": "my-gopher.png"
        },
        "url": {
          "type": "string",
          "example": "https://raw.githubusercontent.com/scraly/gophers/main/arrow-gopher.png"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "HTTP server that handle cute Gophers.",
    "title": "gophers-api",
    "version": "0.1.0"
  },
  "paths": {
    "/gopher": {
      "get": {
        "description": "Get a gopher by a given name",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "Gopher name",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A gopher",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Gopher"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "summary": "Add a new Gopher",
        "parameters": [
          {
            "description": "The Gopher to create.",
            "name": "gopher",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name",
                "path",
                "url"
              ],
              "properties": {
                "name": {
                  "type": "string"
                },
                "path": {
                  "type": "string"
                },
                "url": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Gopher"
            }
          }
        }
      }
    },
    "/gophers": {
      "get": {
        "description": "List Gophers",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "Return the Gophers list.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Gopher"
              }
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "checkHealth",
        "responses": {
          "200": {
            "description": "OK message.",
            "schema": {
              "type": "string",
              "enum": [
                "OK"
              ]
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Gopher": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "example": "my-gopher"
        },
        "path": {
          "type": "string",
          "example": "my-gopher.png"
        },
        "url": {
          "type": "string",
          "example": "https://raw.githubusercontent.com/scraly/gophers/main/arrow-gopher.png"
        }
      }
    }
  }
}`))
}
