package log

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"
	"rucksack/dir"
	"runtime"
)

var logger *slog.Logger

func Init() {
	file, err := os.OpenFile(path.Join(dir.Get(), "rucksack.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}

	logger = slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
	}))
}

func Error(msg string, args ...any) {
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		tmp := []any{"source", fmt.Sprintf("%s:%d", file, lineNumber)}
		args = append(tmp, args...)
	}

	logger.Error(msg, args...)
}
