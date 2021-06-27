package recieve

import "encoding/json"

type InternalMessage struct {
	ID       string          `json:"id"`
	GID      string          `json:"gid"`
	CID      string          `json:"sid"`
	Expires  uint64          `json:"expires"`
	Features []string        `json:"features"`
	Locale   string          `json:"locale"`
	Domain   string          `json:"domain"`
	Event    string          `json:"event"`
	Data     json.RawMessage `json:"data"`
	Saga     Saga            `json:"saga"`
	Src      Src             `json:"src"`
	ReplyTo  ReplyTo         `json:"replyTo"`
	Tokens   Tokens          `json:"tokens"`
}

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
	ID json.RawMessage `json:"id"`
}

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
