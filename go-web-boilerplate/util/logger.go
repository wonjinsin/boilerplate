package util

import (
	"encoding/json"
	"path/filepath"

	"github.com/juju/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger ...
type Logger struct {
	*zap.SugaredLogger
}

func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(filepath.Base(caller.FullPath()))
}

// NewLogger ...
func NewLogger() (*Logger, error) {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
	  	"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "levelKey": "level",
		  "levelEncoder": "capital",
		  "timeKey": "time",
		  "timeEncoder": "iso8601",
		  "nameKey": "trid",
		  "messageKey": "msg",
		  "callerKey": "caller",
		  "callerEncoder": "short"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, errors.Annotate(err, "logger build failed")
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, errors.Annotate(err, "logger build failed")
	}

	defer logger.Sync()
	sugar := logger.Sugar()

	return &Logger{sugar}, nil
}
