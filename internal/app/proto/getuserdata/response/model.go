package response

import "encoding/json"

type Src struct {
	Name    string  `json:"id"`
	Version Version `json:"version"`
}

type Version struct {
	Major uint8 `json:"major"`
	Minor uint8 `json:"minor"`
	Patch uint8 `json:"patch"`
}

type Data struct {
	Username json.RawMessage `json:"username"`
}

type Error struct {
	Code    int16     `json:"code"`
	Message string    `json:"message"`
	Data    ErrorData `json:"data"`
}

type ErrorData struct {
}

type Errors []Error

type Saga struct {
	SID     string `json:"sid"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Expires uint64 `json:"expires"`
}

type ReplyTo struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routingKey"`
	Event      string `json:"event"`
}

type Tokens struct {
	ID     string `json:"id"`
	Access string `json:"access"`
}
