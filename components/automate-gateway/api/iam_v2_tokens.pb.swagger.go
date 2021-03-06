package api

func init() {
	Swagger.Add("iam_v2_tokens", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/iam/v2/tokens.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/iam/v2/tokens": {
      "get": {
        "operationId": "ListTokens",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListTokensResp"
            }
          }
        },
        "tags": [
          "Tokens"
        ]
      },
      "post": {
        "operationId": "CreateToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateTokenReq"
            }
          }
        ],
        "tags": [
          "Tokens"
        ]
      }
    },
    "/iam/v2/tokens/{id}": {
      "get": {
        "operationId": "GetToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Tokens"
        ]
      },
      "delete": {
        "operationId": "DeleteToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Tokens"
        ]
      },
      "put": {
        "operationId": "UpdateToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID can't be changed; ID used to discover token",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateTokenReq"
            }
          }
        ],
        "tags": [
          "Tokens"
        ]
      }
    },
    "/iam/v2beta/tokens": {
      "get": {
        "operationId": "ListTokens2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.ListTokensResp"
            }
          }
        },
        "tags": [
          "Tokens"
        ]
      },
      "post": {
        "operationId": "CreateToken2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.CreateTokenReq"
            }
          }
        ],
        "tags": [
          "Tokens"
        ]
      }
    },
    "/iam/v2beta/tokens/{id}": {
      "get": {
        "operationId": "GetToken2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.GetTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Tokens"
        ]
      },
      "delete": {
        "operationId": "DeleteToken2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.DeleteTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Tokens"
        ]
      },
      "put": {
        "operationId": "UpdateToken2",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID can't be changed; ID used to discover token",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2.UpdateTokenReq"
            }
          }
        ],
        "tags": [
          "Tokens"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.iam.v2.CreateTokenReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "active": {
          "type": "boolean",
          "format": "boolean"
        },
        "value": {
          "type": "string"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.CreateTokenResp": {
      "type": "object",
      "properties": {
        "token": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Token"
        }
      }
    },
    "chef.automate.api.iam.v2.DeleteTokenResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.GetTokenResp": {
      "type": "object",
      "properties": {
        "token": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Token"
        }
      }
    },
    "chef.automate.api.iam.v2.ListTokensResp": {
      "type": "object",
      "properties": {
        "tokens": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2.Token"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.ResetAllTokenProjectsResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2.Token": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "value": {
          "type": "string"
        },
        "active": {
          "type": "boolean",
          "format": "boolean"
        },
        "created_at": {
          "type": "string"
        },
        "updated_at": {
          "type": "string"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateTokenReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID can't be changed; ID used to discover token"
        },
        "name": {
          "type": "string"
        },
        "active": {
          "type": "boolean",
          "format": "boolean"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2.UpdateTokenResp": {
      "type": "object",
      "properties": {
        "token": {
          "$ref": "#/definitions/chef.automate.api.iam.v2.Token"
        }
      }
    }
  }
}
`)
}
