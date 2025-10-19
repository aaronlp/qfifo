package cmd

import (
	"fmt"

	"github.com/aaronlp/qfifo/queues"
)

func ListQueues() {
	for _, queueName := range queues.GetQueueNames() {
		fmt.Println(queueName)
	}
}
