package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DataLocation            string
	LockRetryMilliseconds   uint64
	LockTimeoutMilliseconds uint64
	Env                     string
}

func Load() Config {
	_ = godotenv.Load()

	retryMs, err := strconv.ParseUint(getEnv("LOCK_RETRY_MS", "10"), 10, 64)
	if err != nil {
		log.Fatal("LOCK_RETRY_MS must be a positive integer")
	}

	timeoutMs, err := strconv.ParseUint(getEnv("LOCK_TIMEOUT_MS", "1000"), 10, 64)
	if err != nil {
		log.Fatal("LOCK_TIMEOUT_MS must be a positive integer")
	}

	cfg := Config{
		DataLocation:            getEnv("DATA_LOCATION", "./data"),
		LockRetryMilliseconds:   retryMs,
		LockTimeoutMilliseconds: timeoutMs,
		Env:                     getEnv("ENV", "production"),
	}

	return cfg
}

// getEnv returns the environment variable or a default value if not set.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

// EnsureDirExists checks if the path exists, and creates it if it doesn't.
func EnsureDirExists(path string) {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755) // 0755 is standard permissions
		if err != nil {
			log.Fatalf("Failed to create directory %s: %v", path, err)
		}
		log.Printf("Created directory: %s", path)
	} else if err != nil {
		log.Fatalf("Error checking directory %s: %v", path, err)
	} else if !info.IsDir() {
		log.Fatalf("%s exists but is not a directory", path)
	}
}
