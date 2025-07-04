package benchmark

import (
	"io"
	"log/slog"

	"github.com/kongsakchai/paint"
)

func newSlogText(fields ...slog.Attr) *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelInfo}).WithAttrs(fields))
}

func newCustomSlogText(fields ...slog.Attr) *slog.Logger {
	return slog.New(paint.NewTextHandler(io.Discard, &paint.HandlerOptions{Level: slog.LevelInfo}).WithAttrs(fields))
}

func newDisabledSlogText(fields ...slog.Attr) *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError}).WithAttrs(fields))
}

func newDisabledCustomSlogText(fields ...slog.Attr) *slog.Logger {
	return slog.New(paint.NewTextHandler(io.Discard, &paint.HandlerOptions{Level: slog.LevelError}).WithAttrs(fields))
}

func fakeSlogFields() []slog.Attr {
	return []slog.Attr{
		slog.Int("int", _tenInts[0]),
		slog.Any("ints", _tenInts),
		slog.String("string", _tenStrings[0]),
		slog.Any("strings", _tenStrings),
		slog.Time("time", _tenTimes[0]),
		slog.Any("times", _tenTimes),
		slog.Any("user1", _oneUser),
		slog.Any("user2", _oneUser),
		slog.Any("users", _tenUsers),
		slog.Any("error", errExample),
	}
}

func fakeSlogArgs() []any {
	return []any{
		"int", _tenInts[0],
		"ints", _tenInts,
		"string", _tenStrings[0],
		"strings", _tenStrings,
		"time", _tenTimes[0],
		"times", _tenTimes,
		"user1", _oneUser,
		"user2", _oneUser,
		"users", _tenUsers,
		"error", errExample,
	}
}
