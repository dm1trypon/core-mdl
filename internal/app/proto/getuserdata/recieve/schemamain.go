package recieve

const SchemaMain = `
{
    "type": "object",
    "required": [
        "id",
        "gid",
        "cid",
        "expires",
        "features",
        "locale",
        "domain",
        "event",
        "data",
        "saga",
        "src",
        "replyTo",
        "tokens"
    ],
    "properties": {
        "id": {
            "type": "string"
        },
        "gid": {
            "type": "string"
        },
        "cid": {
            "type": "string"
        },
        "expires": {
            "type": "integer"
        },
        "features": {
            "type": "array",
            "additionalItems": true,
            "items": {
                "anyOf": [
                    {
                        "type": "string"
                    }
                ]
            }
        },
        "locale": {
            "type": "string"
        },
        "domain": {
            "type": "string"
        },
        "event": {
            "type": "string",
            "enum": [
				"getUserData"
			]
        },
        "data": {
            "type": "object"
        },
        "saga": {
            "type": "object",
            "required": [
                "sid",
                "name",
                "state",
                "expires"
            ],
            "properties": {
                "sid": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "expires": {
                    "type": "integer"
                }
            },
            "additionalProperties": true
        },
        "src": {
            "type": "object",
            "required": [
                "name",
                "version"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "version": {
                    "type": "object",
                    "required": [
                        "major",
                        "minor",
                        "patch"
                    ],
                    "properties": {
                        "major": {
                            "type": "integer"
                        },
                        "minor": {
                            "type": "integer"
                        },
                        "patch": {
                            "type": "integer"
                        }
                    },
                    "additionalProperties": true
                }
            },
            "additionalProperties": true
        },
        "replyTo": {
            "type": "object",
            "required": [
                "exchange",
                "routingKey",
                "event"
            ],
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "routingKey": {
                    "type": "string"
                },
                "event": {
                    "type": "string"
                }
            },
            "additionalProperties": true
        },
        "tokens": {
            "type": "object",
            "required": [
                "id",
                "access"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "access": {
                    "type": "string"
                }
            },
            "additionalProperties": true
        }
    },
    "additionalProperties": true
}
`
