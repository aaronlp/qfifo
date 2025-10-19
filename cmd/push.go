package cmd

import (
	"fmt"

	"github.com/aaronlp/qfifo/queues"
)

func Push(queueName, data string) {
	queueItemNumber, err := queues.Push(queueName, []byte(data))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(queueItemNumber)
}
