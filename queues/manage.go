package queues

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateQueue(name string) (bool, error) {
	if QueueExists(name) {
		return false, errors.New("Queue already exists. Not created.")
	}

	err := os.MkdirAll(buildQueuePath(name), 0755)

	if err != nil && !os.IsExist(err) {
		return false, errors.New("Unable to create queue.")
	}

	err = os.MkdirAll(buildQueueDataPath(name), 0755)

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

// GetQueuesNames returns an array of all current queue names
func GetQueueNames() []string {
	files, err := os.ReadDir(queueLocation)

	if err != nil {
		log.Fatal(err)
	}

	list := make([]string, len(files))

	for i, file := range files {
		list[i] = file.Name()
	}

	return list
}

func getQueueCurrentOffset(queueName string) (int, error) {
	currentFileName := buildQueueDataPath(queueName) + "/current"

	data, err := os.ReadFile(currentFileName)
	if err != nil {
		return 0, err
	}

	str := strings.TrimSpace(string(data))

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

	str := strings.TrimSpace(string(data))

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

	return last - current, nil
}

// GetQueueStats returns the current size, starting offset and ending offset of the queue
func GetQueueStats(queueName string) (int, int, int, error) {
	size, err := GetQueueSize(queueName)
	if err != nil {
		return -1, -1, -1, err
	}

	current, err := getQueueCurrentOffset(queueName)
	if err != nil {
		return -1, -1, -1, err
	}

	last, err := getQueueLastOffset(queueName)
	if err != nil {
		return -1, -1, -1, err
	}

	return size, current, last, nil
}
