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
	"io"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDefaultZapLogger is the recommended way to create a new logger, you could call this function to initialize the root
// logger of your application, and provide it to your components, by fx or manually.
func NewDefaultZapLogger() (logr.Logger, error) {
	// change the configuration in the future if needed.
	logLevel := os.Getenv("LOG_LEVEL")
	var options []zap.Option
	if logLevel != "" {
		parsedLevel, err := zapcore.ParseLevel(logLevel)
		if err != nil {
			return logr.Discard(), err
		}
		options = append(options, zap.IncreaseLevel(parsedLevel))
	}

	options = append(options, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return LargeMessageSkippingCore{core: core, maxSize: 250}
	}))

	zapLogger, err := zap.NewDevelopment(options...)
	if err != nil {
		return logr.Discard(), err
	}

	logger := zapr.NewLogger(zapLogger)
	return logger, nil
}

type LargeMessageSkippingCore struct {
	core    zapcore.Core
	maxSize int
}

func (c LargeMessageSkippingCore) Enabled(level zapcore.Level) bool {
	return c.core.Enabled(level)
}
func (c LargeMessageSkippingCore) With(fields []zapcore.Field) zapcore.Core {
	return c.core.With(fields)
}
func (c LargeMessageSkippingCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return ce.AddCore(entry, c)
}
func (c LargeMessageSkippingCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	return c.core.Write(entry, fields)
}
func (c LargeMessageSkippingCore) Sync() error {
	return c.core.Sync()
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
