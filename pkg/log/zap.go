// Copyright 2022 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package log

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDefaultZapLogger is the recommended way to create a new logger, you could call this function to initialize the root
// logger of your application, and provide it to your components, by fx or manually.
func NewDefaultZapLogger() (logr.Logger, error) {
	// change the configuration in the future if needed.
	envLevel := os.Getenv("LOG_LEVEL")
	logLevel := zap.InfoLevel
	if envLevel != "" {
		var err error
		logLevel, err = zapcore.ParseLevel(envLevel)
		if err != nil {
			return logr.Discard(), err
		}
	}

	envMaxFieldSize := os.Getenv("LOG_MAX_FIELD_SIZE")
	maxFieldSize := 16000
	if envMaxFieldSize != "" {
		var err error
		maxFieldSize, err = strconv.Atoi(envMaxFieldSize)
		if err != nil {
			return logr.Discard(), err
		}
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "@timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		NewReflectedEncoder: func(w io.Writer) zapcore.ReflectedEncoder {
			enc := json.NewEncoder(newTruncatingWriter(w, maxFieldSize))
			enc.SetEscapeHTML(false)
			return enc
		},
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	config.Development = true

	zapLogger, err := config.Build(zap.IncreaseLevel(logLevel))
	if err != nil {
		return logr.Discard(), err
	}

	logger := zapr.NewLogger(zapLogger)
	return logger, nil
}

type truncatingWriter struct {
	writer    io.Writer
	enc       *json.Encoder
	maxLength int
}

func newTruncatingWriter(writer io.Writer, maxLength int) truncatingWriter {
	enc := json.NewEncoder(writer)
	enc.SetEscapeHTML(false)

	return truncatingWriter{
		writer:    writer,
		maxLength: maxLength,
		enc:       enc,
	}
}

func (tr truncatingWriter) Write(bytes []byte) (int, error) {
	if len(bytes) > tr.maxLength {
		// Truncating a JSON object most likely yields invalid JSON
		// Passing the truncated string through the JSON encoder yields a valid JSON string
		// This preserves the validity of the overall JSON log entry, while including as much detail as possible of the field being encoded
		return 0, tr.enc.Encode(string(append([]byte("TRUNCATED "), bytes[:tr.maxLength]...)))
	}

	return tr.writer.Write(bytes)
}

// NewZapLoggerWithWriter creates a new logger with io.writer
// The provided encoder presets NewDevelopmentEncoderConfig used by NewDevelopmentConfig do not enable function name logging.
// To enable function name, a non-empty value for config.EncoderConfig.FunctionKey.
func NewZapLoggerWithWriter(out io.Writer) logr.Logger {
	bWriter := out
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.FunctionKey = "function"
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig), zapcore.AddSync(bWriter), config.Level)
	zapLogger := zap.New(core)
	logger := zapr.NewLogger(zapLogger)
	return logger
}
