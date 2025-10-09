package queues

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const queueLocation = "data/queues"
const queueDataLocation = "data/queueData"

func CreateQueue(name string) (bool, error) {
	if QueueExists(name) {
		return false, errors.New("Queue already exists. Not created.")
	}

	err := os.MkdirAll(buildQueuePath(name), 0750)

	if err != nil && !os.IsExist(err) {
		return false, errors.New("Unable to create queue.")
	}

	err = os.MkdirAll(buildQueueDataPath(name), 0750)

	if err != nil && !os.IsExist(err) {
		return false, errors.New("Unable to create queue.")
	}

	currentFileName := buildQueueDataPath(name) + "/current"
	lastFileName := buildQueueDataPath(name) + "/last"

	err = os.WriteFile(currentFileName, []byte("0"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return false, errors.New("Unable to create queue.")
	}

	err = os.WriteFile(lastFileName, []byte("0"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return false, errors.New("Unable to create queue.")
	}

	return true, nil
}

func QueueExists(queueName string) bool {
	dirPath := buildQueuePath(queueName)
	info, _ := os.Stat(dirPath)

	return info != nil && info.IsDir()
}

func getQueuesPath() string {
	return strings.Join([]string{".", queueLocation}, "/")
}

func buildQueuePath(queueName string) string {
	return strings.Join([]string{getQueuesPath(), queueName}, "/")
}

func getQueuesDataPath() string {
	return strings.Join([]string{".", queueDataLocation}, "/")
}

func buildQueueDataPath(queueName string) string {
	return strings.Join([]string{getQueuesDataPath(), queueName}, "/")
}

func ListQueues() {
	files, err := os.ReadDir(getQueuesPath())

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Info())
	}
}

func getQueueCurrentOffset(queueName string) (int, error) {
	currentFileName := buildQueueDataPath(queueName) + "/current"

	data, err := os.ReadFile(currentFileName)
	if err != nil {
		return 0, err
	}

	// Remove any whitespace/newlines
	str := strings.TrimSpace(string(data))

	// Convert string to int
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func getQueueLastOffset(queueName string) (int, error) {
	currentFileName := buildQueueDataPath(queueName) + "/last"

	data, err := os.ReadFile(currentFileName)
	if err != nil {
		return 0, err
	}

	// Remove any whitespace/newlines
	str := strings.TrimSpace(string(data))

	// Convert string to int
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func GetQueueSize(queueName string) (int, error) {
	current, err := getQueueCurrentOffset(queueName)
	if err != nil {
		return 0, err
	}

	last, err := getQueueLastOffset(queueName)
	if err != nil {
		return 0, err
	}

	return (last - current) + 1, nil
}
