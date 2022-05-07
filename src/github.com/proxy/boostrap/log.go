package boostrap

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	logrusgce "github.com/Andreson/logrus-gce"
	log "github.com/sirupsen/logrus"
)

func LogConfig() {
	log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,

		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			pathArray := strings.Split(f.File, "/")
			file = " - " + pathArray[len(pathArray)-1]

			funcArray := strings.Split(f.Function, "/")
			function = " : " + funcArray[len(funcArray)-1] + " : "

			return
		},
		PadLevelText: true,
	})
	log.SetFormatter(logrusgce.NewGCEFormatter(true))
	log.SetLevel(APP_CONFIG.LogLevel)
	Log = log.WithFields(log.Fields{"cache-provider": APP_CONFIG.Provider})
	log.Info("Level Log application : ", log.GetLevel())

}

type MyJSONFormatter struct {
}

func (f *MyJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Note this doesn't include Time, Level and Message which are available on
	// the Entry. Consult `godoc` on information about those fields or read the
	// source of the official loggers.

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %w", err)
	}
	return append(serialized, '\n'), nil
}
