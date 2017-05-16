package log

import (
	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

// Logger はloggerのモデル
type Logger struct {
	logger  *logrus.Logger
	isDebug bool
}

// NewLogger はLoggerインスタンスを生成する
func NewLogger() *Logger {
	logger := &Logger{logger: logrus.New()}
	logger.logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	logger.logger.Out = colorable.NewColorableStdout()
	return logger
}

// NewDebugLogger はデバッグ用のLoggerインスタンスを生成する
func NewDebugLogger() *Logger {
	logger := NewLogger()
	logger.isDebug = true
	return logger
}

// Debug はdebugレベルのログを出力する
func (logger *Logger) Debug(v ...interface{}) {
	logger.logger.Debug(v)
}

// Info はinfoレベルのログを出力する
func (logger *Logger) Info(v ...interface{}) {
	logger.logger.Info(v)
}

// Warn はwarnレベルのログを出力する
func (logger *Logger) Warn(v ...interface{}) {
	logger.logger.Warn(v)
}

// Error はerrorレベルのログを出力する
func (logger *Logger) Error(v ...interface{}) {
	logger.logger.Error(v)
}

// Fatal はfatalレベルのログを出力する
func (logger *Logger) Fatal(v ...interface{}) {
	logger.logger.Fatal(v)
}
