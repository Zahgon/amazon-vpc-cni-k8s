// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package logger

import (
	"github.com/go-logr/logr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type structuredLogger struct {
	zapLogger *zap.SugaredLogger
}

// getZapLevel converts log level string to zapcore.Level.
func getZapLevel(inputLogLevel string) zapcore.Level {
	_ = "STUB: not implemented"
	return *new(zapcore.Level)
}

func (logf *structuredLogger) Debugf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (logf *structuredLogger) Debug(format string) { _ = "STUB: not implemented"; return }

func (logf *structuredLogger) Infof(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (logf *structuredLogger) Info(format string) { _ = "STUB: not implemented"; return }

func (logf *structuredLogger) Warnf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (logf *structuredLogger) Warn(format string) { _ = "STUB: not implemented"; return }

func (logf *structuredLogger) Error(format string) { _ = "STUB: not implemented"; return }

func (logf *structuredLogger) Errorf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (logf *structuredLogger) Fatalf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (logf *structuredLogger) Panicf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (logf *structuredLogger) WithFields(fields Fields) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func getEncoder() zapcore.Encoder { _ = "STUB: not implemented"; return *new(zapcore.Encoder) }

// createZapLogger creates a zap.Logger with the given configuration and caller skip.
func (logConfig *Configuration) createZapLogger(callerSkip int) *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

func (logConfig *Configuration) newZapLogger() *structuredLogger {
	_ = "STUB: not implemented"
	return nil
}

// getPluginLogFilePath returns the writer.
func getPluginLogFilePath(logFilePath string) zapcore.WriteSyncer {
	_ = "STUB: not implemented"
	return *new(zapcore.WriteSyncer)
}

// When path is explicitly empty, write to stderr

// getLogWriter is for lumberjack.
func getLogWriter(logFilePath string) zapcore.WriteSyncer {
	_ = "STUB: not implemented"
	return *new(zapcore.WriteSyncer)
}

// DefaultLogger creates and returns a new default logger.
func DefaultLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

// NewControllerRuntimeLogger creates a logr.Logger compatible with controller-runtime.
func (logConfig *Configuration) NewControllerRuntimeLogger() logr.Logger {
	_ = "STUB: not implemented"
	return *new(logr.Logger)
}
