package queues

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Push adds an item the the provided queue and returns the position in the queue of the item.
func Push(queueName string, data []byte) (int, error) {
	lockFile, err := lock(queueName)
	if err != nil {
		return 0, err
	}

	defer unlock(lockFile)

	offset, err := getQueueLastOffset(queueName)
	if err != nil {
		return 0, err
	}

	offset++

	lastFileName := buildQueueDataPath(queueName) + "/last"
	queueItemFileName := buildQueuePath(queueName) + "/" + fmt.Sprintf("%d", offset)

	err = os.WriteFile(queueItemFileName, data, 0644)
	if err != nil {
		return 0, errors.New("Unable to create item on queue")
	}

	err = os.WriteFile(lastFileName, []byte(fmt.Sprintf("%d", offset)), 0644)
	if err != nil {
		return 0, errors.New("Unable to increment queue counter")
	}

	return offset, nil
}

// Pop retrieves the next item off the queue.
func Pop(queueName string) ([]byte, error) {
	lockFile, err := lock(queueName)
	if err != nil {
		return nil, err
	}

	defer unlock(lockFile)

	size, current, _, err := GetQueueStats(queueName)

	if err != nil {
		return nil, err
	}

	// queue empty
	if size < 1 {
		return nil, nil
	}

	current++

	queueItemFileName := buildQueuePath(queueName) + "/" + fmt.Sprintf("%d", current)

	data, err := os.ReadFile(queueItemFileName)

	if err != nil {
		return nil, err
	}

	currentFileName := buildQueueDataPath(queueName) + "/current"

	err = os.WriteFile(currentFileName, []byte(fmt.Sprintf("%d", current)), 0644)
	if err != nil {
		return nil, errors.New("Unable to increment queue counter")
	}

	return data, nil
}

func lock(queueName string) (*os.File, error) {
	path := queueDataLocation + "/.lock." + queueName
	deadline := time.Now().Add(time.Duration(localConfig.LockTimeoutMilliseconds * uint64(time.Millisecond)))

	for {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0644)

		if err == nil {
			return file, nil
		}

		if os.IsExist(err) {
			if time.Now().After(deadline) {
				return nil, errors.New("Unable to acquire lock")
			}

			time.Sleep(time.Duration(localConfig.LockRetryMilliseconds * uint64(time.Millisecond)))
			continue
		}

		return nil, errors.New("Unknown error, lock not acquired but lock file doesn't exist either")
	}
}

func unlock(lockFile *os.File) error {
	err := lockFile.Close()
	os.Remove(lockFile.Name())
	return err
}
