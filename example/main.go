package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/kongsakchai/paint"
)

type address struct {
	City  string
	State string
}

type user struct {
	id   int
	Name string
	Age  int
	Time time.Time
	Addr *address
}

var users = []user{
	{id: 0, Name: "Alice", Age: 30, Addr: &address{City: "Wonderland", State: "Fantasy"}},
	{id: 1, Name: "Bob", Age: 25, Addr: &address{City: "Builderland", State: "Construction"}},
	{id: 2, Name: "Charlie", Age: 35, Addr: nil},
}

var preAttrs = slog.String("pre", "attributes")

var m = map[string]interface{}{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
}

func main() {
	d := slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	))

	print(d)

	l := slog.New(paint.NewTextHandler(
		os.Stdout,
		&paint.HandlerOptions{
			Level:      slog.LevelDebug,
			TimeFormat: "2006-01-02 15:04:05",
		},
	))

	print(l)
}

func print(l *slog.Logger) {
	l.Debug("Debug message", slog.String("key", "value"))
	l.Info("Info message", slog.String("key", "value"))
	l.Warn("Warning message", slog.String("key", "value"))
	l.Error("Error message", slog.String("key", "value"))

	lw := l.With(slog.String("pre", "attributes"))
	lw.Info("Info message with pre attributes", slog.String("key", "value"))

	lg := l.WithGroup("group")
	lg.Info("Info message with group", slog.String("key", "value"))

	l.Info("Info message and \"any\" data",
		"bool", true,
		"time", time.Now(),
		slog.Any("timeAttr", time.Now()),
		slog.Any("array", []string{
			"January",
			"February",
			"March",
			"April",
			"May",
			"June",
			"July",
			"August",
			"September",
			"October",
			"November",
			"December",
		}),
		slog.Any("odd", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}),
		"even", []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
		slog.Group("users",
			slog.String("type", "json"),
			slog.Any("data", users),
		),
		"map", m,
	)

	fmt.Println("----------------------------")
}
