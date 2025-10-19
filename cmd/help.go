package cmd

import "fmt"

func Help() {
	fmt.Println("create-queue <queueName> - create a new queue\n")
	fmt.Println("list-queues - display a list of all current queues\n")
	fmt.Println("queue-stats <queueName> - display stats for the queue. Format: <queueName> <currentSize> <currentOffset> <lastOffset>\n")
	fmt.Println("push <queueName> <data> - add the provided data onto the queue\n")
	fmt.Println("pop <queueName> - get the next item off the queue and print it\n")
}
