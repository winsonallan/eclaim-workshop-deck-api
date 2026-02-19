package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(env string) (*zap.Logger, error) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Daily rotating file writer
	// Produces: logs/app-2024-01-10.log, logs/app-2024-01-11.log, etc.
	rotator, err := rotatelogs.New(
		filepath.Join("logs", "app-%Y-%m-%d.log"),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(30*24*time.Hour),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize log rotator: %w", err)
	}

	var level zapcore.Level
	if env == "production" {
		level = zapcore.InfoLevel
	} else {
		level = zapcore.DebugLevel
	}

	core := zapcore.NewTee(
		// File — always JSON
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.AddSync(rotator),
			level,
		),
		// Terminal — colored in dev, JSON in prod
		zapcore.NewCore(
			func() zapcore.Encoder {
				if env == "production" {
					return zapcore.NewJSONEncoder(encoderCfg)
				}
				consoleCfg := encoderCfg
				consoleCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
				return zapcore.NewConsoleEncoder(consoleCfg)
			}(),
			zapcore.AddSync(os.Stdout),
			level,
		),
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)), nil
}
