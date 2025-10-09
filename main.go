package main

import (
	"fmt"
	"os"

	"github.com/aaronlp/qfifo/queues"
)

func main() {
	args := os.Args[1:]

	if args[0] == "list" {
		queues.ListQueues()
		os.Exit(0)
	}

	if args[0] == "create" && len(args) == 2 {
		_, err := queues.CreateQueue(args[1])

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Queue was created")
		os.Exit(0)
	}

	if args[0] == "stats" && len(args) == 2 {
		size, err := queues.GetQueueSize(args[1])

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(size)
		os.Exit(0)
	}

	fmt.Println("Options are 'list' or 'create <queueName>'")
	os.Exit(1)
}
