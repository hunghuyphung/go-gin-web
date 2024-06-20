package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLog() {
	log.SetLevel(getLogLevel(os.Getenv("LOG_LEVEL")))
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})

}

func getLogLevel(value string) log.Level {
	switch value {
	case "debug":
		return log.DebugLevel
	case "trace":
		return log.TraceLevel
	default:
		return log.InfoLevel
	}
}
