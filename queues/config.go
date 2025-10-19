package queues

import (
	"github.com/aaronlp/qfifo/internal/config"
	"strings"
)

var queueLocation string
var queueDataLocation string
var localConfig config.Config

func Init(cfg config.Config) {
	config.EnsureDirExists(cfg.DataLocation)

	queueLocation = cfg.DataLocation + "/queues"
	queueDataLocation = cfg.DataLocation + "/queueData"

	localConfig = cfg
}

func buildQueuePath(queueName string) string {
	return strings.Join([]string{queueLocation, queueName}, "/")
}

func buildQueueDataPath(queueName string) string {
	return strings.Join([]string{queueDataLocation, queueName}, "/")
}
