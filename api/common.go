package api

type Payload struct {
	Status  string      `json:"status"`
	Payload interface{} `json:"payload"`
}
