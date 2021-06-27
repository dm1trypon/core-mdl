package response

type IMSuccess struct {
	ID       string   `json:"id"`
	GID      string   `json:"gid"`
	CID      string   `json:"sid"`
	Expires  uint64   `json:"expires"`
	Features []string `json:"features"`
	Locale   string   `json:"locale"`
	Domain   string   `json:"domain"`
	Event    string   `json:"event"`
	Data     Data     `json:"data"`
	Saga     Saga     `json:"saga"`
	Src      Src      `json:"src"`
	ReplyTo  ReplyTo  `json:"replyTo"`
	Tokens   Tokens   `json:"tokens"`
}
