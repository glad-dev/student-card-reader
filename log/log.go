package log

import (
	"log"
	"log/slog"
	"os"
	"path"
	"rucksack/dir"
)

var logger *slog.Logger

func Init() {
	file, err := os.OpenFile(path.Join(dir.Get(), "rucksack.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}

	logger = slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
	}))
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
