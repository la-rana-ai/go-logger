package logger

import (
	"context"
	"github.com/google/uuid"
	"github.com/la-rana-ai/go-logger/mock"
	"log"
	"testing"
)

func createPointerSettingInWhichAllFlagsHaveTheSameStatusForPlain() *OptionFlags {
	return &OptionFlags{
		Date:         false,
		Time:         false,
		Microseconds: false,
		LongFile:     false,
		ShortFile:    true,
		Utc:          false,
		MsgPrefix:    true,
		StdFlags:     false,
	}
}

func Test_tracePlain(t *testing.T) {
	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForPlain())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileDebug := new(mock.File)
	mockFileDebug.On("Write", []byte("plain.go:14: prefixprefix debug Debug message "+uid+" [x-request-id "+uid+"]\n")).Return(136, nil)

	mockFileNoDebug := new(mock.File)

	mockFileInfo := new(mock.File)
	mockFileInfo.On("Write", []byte("plain.go:14: prefixprefix info Info message "+uid+" [x-request-id "+uid+"]\n")).Return(132, nil)

	type args struct {
		logger  *Logger
		ctx     context.Context
		level   InterfaceLevel
		message string
		context []any
	}
	tests := []struct {
		name string
		args args
		mock *mock.File
	}{
		{
			name: "Debug",
			args: args{
				logger: &Logger{
					minimalLevel: Debug,
					flags:        f,
					log:          log.New(mockFileDebug, "prefix", f),
					channel:      "prefix",
				},
				ctx:     ctx,
				level:   Debug,
				message: "Debug message",
				context: []any{
					"x-request-id", ctx.Value("x-request-id"),
				},
			},
			mock: mockFileDebug,
		},
		{
			name: "Nodebug",
			args: args{
				logger: &Logger{
					minimalLevel: Info,
					flags:        f,
					log:          log.New(mockFileNoDebug, "prefix", f),
					channel:      "prefix",
				},
				ctx:     ctx,
				level:   Debug,
				message: "Debug message",
				context: []any{
					"x-request-id", ctx.Value("x-request-id"),
				},
			},
			mock: mockFileNoDebug,
		},
		{
			name: "Info",
			args: args{
				logger: &Logger{
					minimalLevel: Info,
					flags:        f,
					log:          log.New(mockFileInfo, "prefix", f),
					channel:      "prefix",
				},
				ctx:     ctx,
				level:   Info,
				message: "Info message",
				context: []any{
					"x-request-id", ctx.Value("x-request-id"),
				},
			},
			mock: mockFileInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracePlain(tt.args.logger, tt.args.ctx, tt.args.level, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}
