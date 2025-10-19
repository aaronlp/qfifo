package main

import (
	"fmt"
	"os"

	"github.com/aaronlp/qfifo/cmd"
	"github.com/aaronlp/qfifo/internal/config"
	"github.com/aaronlp/qfifo/queues"
)

func main() {
	cfg := config.Load()

	queues.Init(cfg)

	switch os.Args[1] {
	case "list-queues":
		cmd.ListQueues()
	case "create-queue":
		cmd.CreateQueue(os.Args[2])
	case "queue-stats":
		cmd.QueueStats(os.Args[2])
	case "push":
		cmd.Push(os.Args[2], os.Args[3])
	case "pop":
		cmd.Pop(os.Args[2])
	case "help":
		cmd.Help()
	default:
		fmt.Println("Unknown command: ", os.Args[1])
		os.Exit(1)
	}

	os.Exit(0)
}
