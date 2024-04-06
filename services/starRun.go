package services

import (
	"fmt"
	"snake-tail/domain"
	"snake-tail/workers"
)

func StarRun(strRequest string) string {
	var request domain.RequestMessage
	request.Request = strRequest
	request.RepChan = make(chan string)

	workers.ReqChan <- request

	res := <-request.RepChan

	fmt.Println("res: ", res)

	return res
}
