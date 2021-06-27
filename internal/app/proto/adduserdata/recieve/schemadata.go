package recieve

const SchemaData = `
{
	"type": "object",
	"required": [
		"id",
		"nickname"
	],
	"properties": {
		"id": {
			"type": "string"
		},
		"nickname": {
			"type": "string"
		}
	},
	"additionalProperties": true
}
`
