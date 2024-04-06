package domain

type RequestMessage struct {
	Request string
	RepChan chan string
}

type ModelResp struct {
	Resp string `json:"resp"`
}