package main

import (
	"context"
	"github.com/google/uuid"
	"la-rana-ai/go-logger"
)

func main() {

	ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
	log, _ := logger.New("testChannel", &logger.Option{
		Format:       logger.FormatPlain,
		MinimalLevel: logger.Notice,
		Output:       logger.STDOUT,
		Flags: &logger.OptionFlags{
			Date:         true,
			Time:         true,
			Microseconds: true,
			LongFile:     true,
			ShortFile:    false,
			Utc:          true,
			MsgPrefix:    true,
			StdFlags:     true,
		},
	})
	log.Debug(ctx, "этот сообщение не будет записано")
	log.Notice(ctx, "а это сообщение будет записано")
	log.Notice(ctx, "сообщение с контекстом", struct {
		Name string
	}{
		Name: "John Doe",
	})
}
