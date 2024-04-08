package workers

import (
	// "encoding/json"
	"fmt"
	// "bufio"
	// "fmt"
	"io"
	"snake-tail/domain"

	log "github.com/sirupsen/logrus"
)

const (
	// processLabel = "name"
	// errorLabel   = "error"

	ErrorTrue  = "true"
	ErrorFalse = "false"
)

type ErrorResponse struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func Process(worker *domain.Worker, channel chan domain.RequestMessage) {

	for request := range channel {

		// Add response.status = false for each err if the other method is used
		// var response domain.ModelResp
		message := request.Request + "\n"

		// Write message to worker's stdin with retry
		retryCount := 3
		for i := 0; i < retryCount; i++ {
			_, err := io.WriteString(worker.Stdin, message)
			if err != nil {
				if err.Error() == "write |1: broken pipe" {
					log.Warn("Encountered broken pipe error while writing to worker.Stdin. Retrying...")
					continue
				}
				log.Error("Error writing to worker.Stdin: ", err)
				continue
			}
			fmt.Println("Message sent to worker: ", message)
			break
		}

		// Read result from worker's stdout
		result, err := worker.StdOutReader.ReadString('\n')
		if err != nil {
			log.Error("Error reading from worker.StdOutReader: ", err)
			continue
		}

		// // Unmarshal the result into a struct
		// err = json.Unmarshal(result, &response)
		// if err != nil {
		// 	log.Error("Error unmarshalling result: ", err)
		// 	continue
		// }

		snakeName := result[1 : len(result)-2] // Adjust indices as per the actual format of the result

		fmt.Println("Snake name: ", snakeName)

		// Send result back through response channel
		request.RepChan <- snakeName
	}
}
