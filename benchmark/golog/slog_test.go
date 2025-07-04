package benchmark

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/kongsakchai/paint"
)

func slogAttrs() []slog.Attr {
	return []slog.Attr{
		slog.Int("bytes", ctxBodyBytes),
		slog.String("request", ctxRequest),
		slog.Float64("elapsed_time_ms", ctxTimeElapsedMs),
		slog.Any("user", ctxUser),
		slog.Time("now", ctxTime),
		slog.Any("months", ctxMonths),
		slog.Any("primes", ctxFirst10Primes),
		slog.Any("users", ctxUsers),
		slog.Any("error", ctxErr),
	}
}

func newSlogText(w io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func newSlogTextWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}).WithAttrs(attr))
}

type slogBench struct {
	l *slog.Logger
}

func (b *slogBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *slogBench) logEventFmt(msg string, args ...any) {
	b.l.Info(fmt.Sprintf(msg, args...))
}

func (b *slogBench) logEventCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		msg,
		slogAttrs()...,
	)
}

func (b *slogBench) logEventCtxWeak(msg string) {
	b.l.Info(msg, alternatingKeyValuePairs()...)
}

func (b *slogBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *slogBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debug(fmt.Sprintf(msg, args...))
}

func (b *slogBench) logDisabledCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelDebug,
		msg,
		slogAttrs()...,
	)
}

func (b *slogBench) logDisabledCtxWeak(msg string) {
	b.l.Debug(msg, alternatingKeyValuePairs()...)
}

func newCutomSlogText(w io.Writer) *slog.Logger {
	return slog.New(paint.NewTextHandler(w, &paint.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func newCutomSlogTextWithCtx(w io.Writer, attr []slog.Attr) *slog.Logger {
	return slog.New(paint.NewTextHandler(w, &paint.HandlerOptions{
		Level: slog.LevelInfo,
	}).WithAttrs(attr))
}

// customSlogTextBench is a benchmark for the custom slog logger

type slogTextBench struct {
	*slogBench
}

func (b *slogTextBench) new(w io.Writer) logBenchmark {
	return &slogTextBench{
		&slogBench{
			l: newSlogText(w),
		},
	}
}

func (b *slogTextBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogTextBench{
		&slogBench{
			l: newSlogTextWithCtx(w, slogAttrs()),
		},
	}
}

func (b *slogTextBench) name() string {
	return "Slog Text"
}

type customSlogTextBench struct {
	*slogBench
}

func (b *customSlogTextBench) new(w io.Writer) logBenchmark {
	return &customSlogTextBench{
		&slogBench{
			l: newCutomSlogText(w),
		},
	}
}

func (b *customSlogTextBench) newWithCtx(w io.Writer) logBenchmark {
	return &customSlogTextBench{
		&slogBench{
			l: newCutomSlogTextWithCtx(w, slogAttrs()),
		},
	}
}

func (b *customSlogTextBench) name() string {
	return "Custom Slog Text"
}
