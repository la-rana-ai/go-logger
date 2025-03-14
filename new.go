package logger

import (
	"context"
	"io"
	"log"
	"os"
)

type Interface interface {
	Debug(ctx context.Context, message string, context ...any)
	Info(ctx context.Context, message string, context ...any)
	Notice(ctx context.Context, message string, context ...any)
	Warning(ctx context.Context, message string, context ...any)
	Error(ctx context.Context, message string, context ...any)
	Critical(ctx context.Context, message string, context ...any)
	Alert(ctx context.Context, message string, context ...any)
	Emergency(ctx context.Context, message string, context ...any)
}

var channels = make(map[string]Interface)

func New(name string, option InterfaceOption) (Interface, error) {
	if channel, ok := channels[name]; ok {
		return channel, nil
	}
	var (
		flags          int
		flagsLogModule int
		out            io.Writer
		prefix         = ""
		err            error
		trace          func(logger *Logger, ctx context.Context, level InterfaceLevel, message string, context ...any)
		minLevel       InterfaceLevel
	)

	if option != nil {
		flags = getFlag(option.GetFlags())
		switch option.GetFormat().Num() {
		case FormatJSON.Num():
			trace = traceJson
		case FormatPlain.Num(): //
			fallthrough
		default:
			flagsLogModule = flags
			prefix = name + ":"
			trace = tracePlain
		}

		if out, err = getOutput(option.GetOutput(), option.GetFile()); err != nil {
			return nil, err
		}
		minLevel = option.GetMinimalLevel()
		if minLevel == nil {
			minLevel = Debug
		}
	} else {
		out = os.Stdout
		trace = traceJson
		minLevel = Debug
	}

	channels[name] = &Logger{
		minimalLevel: minLevel,
		channel:      name,
		log:          log.New(out, prefix, flagsLogModule),
		trace:        trace,
		flags:        flags,
	}
	return channels[name], nil
}
