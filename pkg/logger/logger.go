package logger

import (
	"os"
	"sync"
	"test/pkg/config"

	"github.com/sirupsen/logrus"
)

const key = "logger"
const logType = "log_type"

// structured data log fields.
const (
	FieldUUID       = "uuid"
	FieldStackTrace = "stack_trace"
)

var (
	// singleton instance
	log *logrus.Logger
)

var once sync.Once

func New(cfg config.Config) *logrus.Logger {


	once.Do(func() { // executes only once
		log = buildLogger(cfg)
	})

	return log
}

func buildLogger(cfg config.Config) *logrus.Logger {
	log := logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log this severity or above.
	log.SetLevel(logrus.InfoLevel)

	// Log file name and line number as "file" and
	// Caller function name as "func"
	log.SetReportCaller(true)

	return log
}