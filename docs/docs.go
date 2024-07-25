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
        "/anime/demoKafka": {
            "get": {
                "description": "Send kafka message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kafka Demo"
                ],
                "summary": "Send kafka message",
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/anime.Anime"
                        }
                    }
                }
            }
        },
        "/anime/{id}": {
            "get": {
                "description": "Get an anime by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Animes"
                ],
                "summary": "Get an anime by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Anime ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/anime.Anime"
                        }
                    }
                }
            }
        },
        "/characters/create": {
            "post": {
                "description": "create character and insert to elasticsearch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Characters"
                ],
                "summary": "create character and insert to elasticsearch",
                "parameters": [
                    {
                        "description": "Character JSON",
                        "name": "inputCharacter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/character.CreateCharacterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/character.Character"
                        }
                    }
                }
            }
        },
        "/elasticsearch/create-character-index": {
            "post": {
                "description": "create elastic character_index",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Elasticsearch"
                ],
                "summary": "create character index for elasticsearch",
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/elasticsearch.CreateElasticIndexResponse"
                        }
                    }
                }
            }
        },
        "/files/uploadFileImage": {
            "post": {
                "description": "Upload a file image and store information",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "summary": "Upload a file image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "File name",
                        "name": "fileName",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "File description",
                        "name": "description",
                        "in": "formData"
                    },
                    {
                        "enum": [
                            "demo/characters",
                            "demo/animes"
                        ],
                        "type": "string",
                        "description": "Information type",
                        "name": "infoType",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Updated by",
                        "name": "updateBy",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/file_info.FileInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "anime.Anime": {
            "type": "object",
            "properties": {
                "activate": {
                    "type": "boolean"
                },
                "createAt": {
                    "type": "string"
                },
                "createBy": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                }
            }
        },
        "character.Character": {
            "type": "object",
            "properties": {
                "activate": {
                    "type": "boolean"
                },
                "animeID": {
                    "type": "string"
                },
                "createAt": {
                    "type": "string"
                },
                "createBy": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                }
            }
        },
        "character.CreateCharacterRequest": {
            "type": "object",
            "properties": {
                "animeID": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                }
            }
        },
        "elasticsearch.CreateElasticIndexResponse": {
            "type": "object",
            "properties": {
                "index": {
                    "type": "string"
                }
            }
        },
        "file_info.FileInfo": {
            "type": "object",
            "properties": {
                "activate": {
                    "type": "boolean"
                },
                "createAt": {
                    "type": "string"
                },
                "createBy": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "extension": {
                    "type": "string"
                },
                "fileName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "infoType": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}