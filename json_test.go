package logger

import (
	"context"
	"github.com/google/uuid"
	"la-rana-ai/go-logger/mock"
	"log"
	"regexp"
	"testing"
)

func Test_getFileAndLine(t *testing.T) {
	type args struct {
		flags int
		pc    int
	}
	type rez struct {
		file string
		line uint
	}
	tests := []struct {
		name string
		args args
		want rez
	}{
		{
			name: "Lshortfile",
			args: args{
				flags: log.Lshortfile,
				pc:    0,
			},
			want: rez{
				file: "json_test.go",
				line: 51,
			},
		},
		{
			name: "zero",
			args: args{
				flags: 0,
				pc:    0,
			},
			want: rez{
				file: "",
				line: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, line := getFileAndLine(tt.args.flags, tt.args.pc)
			if file != tt.want.file {
				t.Errorf("getFileAndLine() file = %v, want %v", file, tt.want.file)
			}
			if line != tt.want.line {
				t.Errorf("getFileAndLine() line = %v, want %v", line, tt.want.line)
			}
		})
	}
}

func Test_itoa(t *testing.T) {
	type args struct {
		val int
		min uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zero",
			args: args{
				val: 0,
				min: 1,
			},
			want: "0",
		},
		{
			name: "zero 2",
			args: args{
				val: 0,
				min: 2,
			},
			want: "00",
		},
		{
			name: "zero4",
			args: args{
				val: 0,
				min: 4,
			},
			want: "0000",
		},
		{
			name: "one",
			args: args{
				val: 1,
				min: 1,
			},
			want: "1",
		},
		{
			name: "one 2",
			args: args{
				val: 1,
				min: 2,
			},
			want: "01",
		},
		{
			name: "one4",
			args: args{
				val: 1,
				min: 4,
			},
			want: "0001",
		},
		{
			name: "ten",
			args: args{
				val: 10,
				min: 1,
			},
			want: "10",
		},
		{
			name: "ten 2",
			args: args{
				val: 10,
				min: 2,
			},
			want: "10",
		},
		{
			name: "ten 4",
			args: args{
				val: 10,
				min: 4,
			},
			want: "0010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itoa(tt.args.val, tt.args.min); got != tt.want {
				t.Errorf("itoa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDatetime(t *testing.T) {
	type want struct {
		isNil    bool
		mask     *regexp.Regexp
		timezone string
	}
	type args struct {
		flags int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "zero",
			args: args{
				flags: 0,
			},
			want: want{
				isNil: true,
			},
		},
		{
			name: "Ldate",
			args: args{
				flags: log.Ldate,
			},
			want: want{
				isNil:    false,
				mask:     regexp.MustCompile(`^\d{4}/\d{2}/\d{2}$`),
				timezone: "",
			},
		},
		{
			name: "Ltime",
			args: args{
				flags: log.Ltime,
			},
			want: want{
				isNil:    false,
				mask:     regexp.MustCompile(`^\d{2}:\d{2}:\d{2}$`),
				timezone: "",
			},
		},
		{
			name: "Lmicroseconds",
			args: args{
				flags: log.Lmicroseconds,
			},
			want: want{
				isNil:    false,
				mask:     regexp.MustCompile(`^\d{2}:\d{2}:\d{2}.\d{6}$`),
				timezone: "",
			},
		},
		{
			name: "LstdFlags",
			args: args{
				flags: log.LstdFlags,
			},
			want: want{
				isNil:    false,
				mask:     regexp.MustCompile(`^\d{4}/\d{2}/\d{2}T\d{2}:\d{2}:\d{2}$`),
				timezone: "",
			},
		},
		{
			name: "LUTC",
			args: args{
				flags: log.LUTC | log.LstdFlags,
			},
			want: want{
				isNil:    false,
				mask:     regexp.MustCompile(`^\d{4}/\d{2}/\d{2}T\d{2}:\d{2}:\d{2}$`),
				timezone: "UTC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDatetime(tt.args.flags)
			if got != nil {
				if tt.want.isNil {
					t.Errorf("getDatetime() = %v, want nil", got)
				}
				if got.Timezone != tt.want.timezone {
					t.Errorf("getDatetime().Timezone = %v, want %v", got.Timezone, tt.want.timezone)
				}
				if !tt.want.mask.MatchString(got.Date) {
					t.Errorf("getDatetime().Date = %v does not match the mask", got.Date)
				}

			} else {
				if !tt.want.isNil {
					t.Errorf("getDatetime() = %v, want not nil", got)
				}
			}
		})
	}
}

func createPointerSettingInWhichAllFlagsHaveTheSameStatusForJson() *OptionFlags {
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
func Test_traceJson(t *testing.T) {
	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForJson())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileDebug := new(mock.File)
	mockFileDebug.On("Write", []byte("{\"channel\":\"prefix\",\"level-name\":\"debug\",\"level\":7,\"file\":\"testing.go\",\"line\":1792,\"message\":\"Debug message\",\"context\":{\"args\":[\"x-request-id\",\""+uid+"\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(314, nil)

	mockFileNoDebug := new(mock.File)

	mockFileInfo := new(mock.File)
	mockFileInfo.On("Write", []byte("{\"channel\":\"prefix\",\"level-name\":\"info\",\"level\":6,\"file\":\"testing.go\",\"line\":1792,\"message\":\"Info message\",\"context\":{\"args\":[\"x-request-id\",\""+uid+"\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(312, nil)

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
					log:          log.New(mockFileDebug, "", 0),
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
					log:          log.New(mockFileNoDebug, "", 0),
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
					trace:        traceJson,
					minimalLevel: Info,
					flags:        f,
					log:          log.New(mockFileInfo, "", 0),
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
			traceJson(tt.args.logger, tt.args.ctx, tt.args.level, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}
