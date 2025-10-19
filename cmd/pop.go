package cmd

import (
	"fmt"

	"github.com/aaronlp/qfifo/queues"
)

func Pop(queueName string) {
	itemData, err := queues.Pop(queueName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if itemData == nil {
		fmt.Println("<Queue Empty>")
		return
	}

	fmt.Println(string(itemData))
}
