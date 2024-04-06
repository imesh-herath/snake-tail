package workers

import (
	"bufio"
	"log"
	"os/exec"
	"snake-tail/domain"
	"strconv"
)

var (
	ReqChan chan domain.RequestMessage
)

func InitWorkers() {
	ReqChan = make(chan domain.RequestMessage, 1)
	numOfWorkers := 1

	for index := 0; index < numOfWorkers; index++ {
		var worker domain.Worker
		worker.ID = index
		pythonScript := "./model/model.py"

		//creating ports for python models
		pythonPromPort := "800" + strconv.FormatInt(int64(index), 10)

		command := exec.Command("python3", pythonScript, pythonPromPort)

		stdIn, err := command.StdinPipe()
		if err != nil {
			log.Fatalf("Failed to create StdinPipe for worker %d: %v", index, err)
		}

		worker.Stdin = stdIn

		stdOut, err := command.StdoutPipe()
		if err != nil {
			log.Fatalf("Failed to create StdoutPipe for worker %d: %v", index, err)
		}

		worker.StdOutReader = bufio.NewReaderSize(stdOut, 1)

		err = command.Start()
		if err != nil {
			log.Fatalf("Failed to start command for worker %d: %v", index, err)
		}

		go Process(&worker, ReqChan)
	}
}
