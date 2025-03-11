package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/go-jedi/posts/config"
	"github.com/natefinch/lumberjack"
)

type Logger struct {
	*slog.Logger
}

func New(cfg config.LoggerConfig) *Logger {
	ho := &slog.HandlerOptions{}

	levelMapping := map[string]slog.Level{
		"debug":   slog.LevelDebug,
		"info.md": slog.LevelInfo,
		"warn":    slog.LevelWarn,
		"error":   slog.LevelError,
	}

	if v, ok := levelMapping[cfg.Level]; ok {
		ho.Level = v
	} else {
		ho.Level = slog.LevelInfo
	}

	ho.AddSource = cfg.AddSource

	var h slog.Handler = slog.NewTextHandler(os.Stdout, ho)

	if cfg.IsJSON {
		h = slog.NewJSONHandler(os.Stdout, ho)
	}

	if cfg.SetFile {
		mw := io.MultiWriter(
			os.Stdout,
			&lumberjack.Logger{
				Filename:   cfg.FileName,
				MaxSize:    cfg.MaxSize,
				MaxBackups: cfg.MaxBackups,
				MaxAge:     cfg.MaxAge,
			},
		)

		if cfg.IsJSON {
			h = slog.NewJSONHandler(mw, ho)
		} else {
			h = slog.NewTextHandler(mw, ho)
		}
	}

	logger := slog.New(h)

	return &Logger{
		logger,
	}
}
