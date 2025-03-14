package logger

import (
	"context"
	"github.com/google/uuid"
	"la-rana-ai/go-logger/mock"
	"log"
	"testing"
)

func TestLogger_Alert(t *testing.T) {

	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForMoreImportantThanTheMinimum := new(mock.File)
	mockFileForMoreImportantThanTheMinimum.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"alert\",\"level\":1,\"file\":\"logger_test.go\",\"line\":93,\"message\":\"test\",\"context\":{\"args\":[\"123\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(293, nil)

	mockFileForMinimumLevel := new(mock.File)
	mockFileForMinimumLevel.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"alert\",\"level\":1,\"file\":\"logger_test.go\",\"line\":93,\"message\":\"test\",\"context\":{\"args\":[\"ert\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(293, nil)

	mockFileForLessImportantThanTheMinimum := new(mock.File)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Alert MoreImportant",
			fields: &Logger{
				minimalLevel: Debug,
				flags:        f,
				log:          log.New(mockFileForMoreImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"123",
				},
			},
			mock: mockFileForMoreImportantThanTheMinimum,
		},
		{
			name: "TestLogger_Alert minimal level",
			fields: &Logger{
				minimalLevel: Alert,
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"ert",
				},
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Alert LessImportant",
			fields: &Logger{
				minimalLevel: Emergency,
				flags:        f,
				log:          log.New(mockFileForLessImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"ert",
				},
			},
			mock: mockFileForLessImportantThanTheMinimum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Alert(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func TestLogger_Critical(t *testing.T) {

	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForMoreImportantThanTheMinimum := new(mock.File)
	mockFileForMoreImportantThanTheMinimum.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"critical\",\"level\":2,\"file\":\"logger_test.go\",\"line\":181,\"message\":\"test\",\"context\":{\"args\":[\"123\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(261, nil)

	mockFileForMinimumLevel := new(mock.File)
	mockFileForMinimumLevel.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"critical\",\"level\":2,\"file\":\"logger_test.go\",\"line\":181,\"message\":\"test\",\"context\":{\"args\":[\"ert\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(261, nil)

	mockFileForLessImportantThanTheMinimum := new(mock.File)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Critical MoreImportant",
			fields: &Logger{
				minimalLevel: Notice,
				flags:        f,
				log:          log.New(mockFileForMoreImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"123",
				},
			},
			mock: mockFileForMoreImportantThanTheMinimum,
		},
		{
			name: "TestLogger_Critical minimal level",
			fields: &Logger{
				minimalLevel: Critical,
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"ert",
				},
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Critical LessImportant",
			fields: &Logger{
				minimalLevel: Alert,
				flags:        f,
				log:          log.New(mockFileForLessImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"ert",
				},
			},
			mock: mockFileForLessImportantThanTheMinimum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Critical(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func TestLogger_Debug(t *testing.T) {

	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForMinimumLevel := new(mock.File)
	mockFileForMinimumLevel.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"debug\",\"level\":7,\"file\":\"logger_test.go\",\"line\":266,\"message\":\"test\",\"context\":{\"args\":[\"Debug text\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(265, nil)

	mockFileForLessImportantThanTheMinimum := new(mock.File)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Debug MoreImportant",
			fields: &Logger{
				minimalLevel: &Level{int: 8, str: "test"},
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"Debug text",
				},
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Debug minimal level",
			fields: &Logger{
				minimalLevel: Debug,
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"Debug text",
				},
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Debug LessImportant",
			fields: &Logger{
				minimalLevel: Alert,
				flags:        f,
				log:          log.New(mockFileForLessImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"Debug text",
				},
			},
			mock: mockFileForLessImportantThanTheMinimum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Debug(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func TestLogger_Emergency(t *testing.T) {
	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForEmergencyTests := new(mock.File)
	mockFileForEmergencyTests.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"emergency\",\"file\":\"logger_test.go\",\"line\":330,\"message\":\"Emergency message\",\"context\":{\"args\":[\"Emergency context\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(279, nil)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Emergency MoreImportant",
			fields: &Logger{
				minimalLevel: Notice,
				flags:        f,
				log:          log.New(mockFileForEmergencyTests, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Emergency message",
				context: []any{
					"Emergency context",
				},
			},
			mock: mockFileForEmergencyTests,
		},
		{
			name: "TestLogger_Emergency minimal level",
			fields: &Logger{
				minimalLevel: Emergency,
				flags:        f,
				log:          log.New(mockFileForEmergencyTests, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Emergency message",
				context: []any{
					"Emergency context",
				},
			},
			mock: mockFileForEmergencyTests,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Emergency(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func TestLogger_Error(t *testing.T) {

	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForMoreImportantThanTheMinimum := new(mock.File)
	mockFileForMoreImportantThanTheMinimum.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"error\",\"level\":3,\"file\":\"logger_test.go\",\"line\":415,\"message\":\"Error message\",\"context\":{\"args\":[\"Error message\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(277, nil)

	mockFileForMinimumLevel := new(mock.File)
	mockFileForMinimumLevel.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"error\",\"level\":3,\"file\":\"logger_test.go\",\"line\":415,\"message\":\"Error message\",\"context\":{\"args\":null,\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(264, nil)

	mockFileForLessImportantThanTheMinimum := new(mock.File)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Error MoreImportant",
			fields: &Logger{
				minimalLevel: Debug,
				flags:        f,
				log:          log.New(mockFileForMoreImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Error message",
				context: []any{
					"Error message",
				},
			},
			mock: mockFileForMoreImportantThanTheMinimum,
		},
		{
			name: "TestLogger_Error minimal level",
			fields: &Logger{
				minimalLevel: Error,
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Error message",
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Error LessImportant",
			fields: &Logger{
				minimalLevel: Alert,
				flags:        f,
				log:          log.New(mockFileForLessImportantThanTheMinimum, "", 0),
				channel:      "Error message",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"Error message",
				},
			},
			mock: mockFileForLessImportantThanTheMinimum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Error(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func TestLogger_Info(t *testing.T) {

	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForMoreImportantThanTheMinimum := new(mock.File)
	mockFileForMoreImportantThanTheMinimum.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"info\",\"level\":6,\"file\":\"logger_test.go\",\"line\":501,\"message\":\"Info text\",\"context\":{\"args\":[\"Info data\",\"Info data 2\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(282, nil)

	mockFileForMinimumLevel := new(mock.File)
	mockFileForMinimumLevel.On("Write", []byte("{\"channel\":\"InfoTest\",\"level-name\":\"info\",\"level\":6,\"file\":\"logger_test.go\",\"line\":501,\"message\":\"Info text\",\"context\":{\"args\":[\"Info text\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(272, nil)

	mockFileForLessImportantThanTheMinimum := new(mock.File)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Info MoreImportant",
			fields: &Logger{
				minimalLevel: Debug,
				flags:        f,
				log:          log.New(mockFileForMoreImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Info text",
				context: []any{
					"Info data",
					"Info data 2",
				},
			},
			mock: mockFileForMoreImportantThanTheMinimum,
		},
		{
			name: "TestLogger_Info minimal level",
			fields: &Logger{
				minimalLevel: Info,
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "InfoTest",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Info text",
				context: []any{
					"Info text",
				},
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Info LessImportant",
			fields: &Logger{
				minimalLevel: Critical,
				flags:        f,
				log:          log.New(mockFileForLessImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Info text",
			},
			mock: mockFileForLessImportantThanTheMinimum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Info(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func TestLogger_Notice(t *testing.T) {

	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForMoreImportantThanTheMinimum := new(mock.File)
	mockFileForMoreImportantThanTheMinimum.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"notice\",\"level\":5,\"file\":\"logger_test.go\",\"line\":586,\"message\":\"test\",\"context\":{\"args\":[\"Notice data\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(267, nil)

	mockFileForMinimumLevel := new(mock.File)
	mockFileForMinimumLevel.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"notice\",\"level\":5,\"file\":\"logger_test.go\",\"line\":586,\"message\":\"Notice test\",\"context\":{\"args\":null,\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(263, nil)

	mockFileForLessImportantThanTheMinimum := new(mock.File)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Notice MoreImportant",
			fields: &Logger{
				minimalLevel: Debug,
				flags:        f,
				log:          log.New(mockFileForMoreImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"Notice data",
				},
			},
			mock: mockFileForMoreImportantThanTheMinimum,
		},
		{
			name: "TestLogger_Notice minimal level",
			fields: &Logger{
				minimalLevel: Notice,
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Notice test",
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Notice LessImportant",
			fields: &Logger{
				minimalLevel: Alert,
				flags:        f,
				log:          log.New(mockFileForLessImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "test",
				context: []any{
					"ert",
				},
			},
			mock: mockFileForLessImportantThanTheMinimum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Notice(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func TestLogger_Warning(t *testing.T) {

	uid := uuid.New().String()
	f := getFlag(createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger())
	ctx := context.WithValue(context.Background(), "x-request-id", uid)

	mockFileForMoreImportantThanTheMinimum := new(mock.File)
	mockFileForMoreImportantThanTheMinimum.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"warning\",\"level\":4,\"file\":\"logger_test.go\",\"line\":671,\"message\":\"Warning test\",\"context\":{\"args\":[\"Warning context\"],\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(280, nil)

	mockFileForMinimumLevel := new(mock.File)
	mockFileForMinimumLevel.On("Write", []byte("{\"channel\":\"test\",\"level-name\":\"warning\",\"level\":4,\"file\":\"logger_test.go\",\"line\":671,\"message\":\"Warning test\",\"context\":{\"args\":null,\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(265, nil)

	mockFileForLessImportantThanTheMinimum := new(mock.File)

	type args struct {
		ctx     context.Context
		message string
		context []any
	}
	tests := []struct {
		name   string
		fields *Logger
		args   args
		mock   *mock.File
	}{
		{
			name: "TestLogger_Warning MoreImportant",
			fields: &Logger{
				minimalLevel: Notice,
				flags:        f,
				log:          log.New(mockFileForMoreImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Warning test",
				context: []any{
					"Warning context",
				},
			},
			mock: mockFileForMoreImportantThanTheMinimum,
		},
		{
			name: "TestLogger_Warning minimal level",
			fields: &Logger{
				minimalLevel: Warning,
				flags:        f,
				log:          log.New(mockFileForMinimumLevel, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Warning test",
			},
			mock: mockFileForMinimumLevel,
		},
		{
			name: "TestLogger_Warning LessImportant",
			fields: &Logger{
				minimalLevel: Alert,
				flags:        f,
				log:          log.New(mockFileForLessImportantThanTheMinimum, "", 0),
				channel:      "test",
				trace:        traceJson,
			},
			args: args{
				ctx:     ctx,
				message: "Warning test",
				context: []any{
					"Warning context",
				},
			},
			mock: mockFileForLessImportantThanTheMinimum,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Warning(tt.args.ctx, tt.args.message, tt.args.context...)
			tt.mock.AssertExpectations(t)
		})
	}
}

func createPointerSettingInWhichAllFlagsHaveTheSameStatusForLogger() *OptionFlags {
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
