package model

type Payload struct {
	Round int `json:"round"`
	Rand  int `json:"rand"`
}

type Message struct {
	Key     string   `json:"key"`
	Value   string   `json:"value"`
	Payload *Payload `json:"payload"`
}
