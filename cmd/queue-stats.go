package cmd

import (
	"fmt"

	"github.com/aaronlp/qfifo/queues"
)

func QueueStats(queueName string) {
	size, current, offset, err := queues.GetQueueStats(queueName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(queueName, size, current, offset)
}
