package cmd

import (
	"fmt"

	"github.com/aaronlp/qfifo/queues"
)

func CreateQueue(queueName string) {
	_, err := queues.CreateQueue(queueName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Queue was created")
}
