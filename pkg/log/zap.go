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
	"strings"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Environment variable to control log format, if set to JSON, structured logging will be used. Default is "console", and to use zap.NewDevelopmentConfig.
const LogFormat = "LOG_FORMAT"

// Environment variable to control logging level, should parse to zapcore.Level. Default is "info".
const LogLevel = "LOG_LEVEL"

// Environment variable to control logging key used for log message in structured logging. Default is defined by zap.NewProductionEncoderConfig.
const LogKeyMessage = "LOG_KEY_MESSAGE"

// Environment variable to control logging key used for log timestamp in structured logging. Default is defined by zap.NewProductionEncoderConfig.
const LogKeyTimestamp = "LOG_KEY_TIMESTAMP"

// Environment variable to control the maximum field size in structured logging before truncating the field value.
const LogMaxFieldSize = "LOG_MAX_FIELD_SIZE"

// Environment variable to control the logging timestamp format used in structured logging. Valid values are "rfc3339", "rfc3339nano" and "epoch".
const LogTimestampFormat = "LOG_TIMESTAMP_FORMAT"

// NewDefaultZapLogger is the recommended way to create a new logger, you could call this function to initialize the root
// logger of your application, and provide it to your components, by fx or manually.
func NewDefaultZapLogger() (logr.Logger, error) {
	envLevel := os.Getenv(LogLevel)
	logLevel := zap.InfoLevel
	if envLevel != "" {
		var err error
		logLevel, err = zapcore.ParseLevel(envLevel)
		if err != nil {
			return logr.Discard(), err
		}
	}

	envFormat := os.Getenv(LogFormat)
	logFormat := "console"
	if envFormat == "json" {
		logFormat = "json"
	}

	var config zap.Config

	if logFormat == "json" {
		encoderConfig := zap.NewProductionEncoderConfig()

		if v := os.Getenv(LogKeyMessage); v != "" {
			encoderConfig.MessageKey = v
		}
		if v := os.Getenv(LogKeyTimestamp); v != "" {
			encoderConfig.TimeKey = v
		}

		if v := os.Getenv(LogTimestampFormat); v != "" {
			if v == "rfc3339" {
				encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
			} else if v == "rfc3339nano" {
				encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
			} else if v == "epoch" {
				encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
			}
		}

		// If configured, truncate the fields to the configured size. This allows for reasonable configuration to prevent extremely
		// long logging lines (e.g. by logging an object with a large collection of events attached), which can cause issues with
		// log ingestion/collection, while guaranteeing that the output remains valid JSON.
		envMaxFieldSize := os.Getenv(LogMaxFieldSize)
		if envMaxFieldSize != "" {
			maxFieldSize, err := strconv.Atoi(envMaxFieldSize)
			if err == nil {
				encoderConfig.NewReflectedEncoder = func(w io.Writer) zapcore.ReflectedEncoder {
					enc := json.NewEncoder(newTruncatingWriter(w, maxFieldSize))
					enc.SetEscapeHTML(false)
					return enc
				}
			}
		}

		config = zap.NewProductionConfig()
		config.EncoderConfig = encoderConfig
	} else {
		config = zap.NewDevelopmentConfig()
	}

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
	output := bytes

	if len(bytes) > tr.maxLength {
		output = append([]byte("TRUNCATED "), bytes[:tr.maxLength]...)
	}

	// Always encode the field value as a JSON string, this ensures that the log field types will not change between logging
	// of truncated and untruncated values, while also ensuring that the log line remains valid JSON.
	return 0, tr.enc.Encode(strings.TrimRight(string(output), "\n"))
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
