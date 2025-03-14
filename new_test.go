package logger

import (
	"context"
	"github.com/google/uuid"
	"la-rana-ai/go-logger/mock"
	"testing"
)

func createPointerSettingInWhichAllFlagsHaveTheSameStatusForLoggerNew() *OptionFlags {
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

func TestNew(t *testing.T) {
	type args struct {
		name   string
		option InterfaceOption
	}

	uid := uuid.New().String()
	ctx := context.WithValue(context.Background(), "x-request-id", uid)
	mockFileForTestJson := new(mock.File)
	mockFileForTestJson.On("Write", []byte("{\"channel\":\"testChannel\",\"level-name\":\"alert\",\"level\":1,\"file\":\"new_test.go\",\"line\":77,\"message\":\"message\",\"context\":{\"args\":null,\"x-request-id\":\""+uid+"\",\"process-id\":\""+ProcessID.String()+"\",\"host\":\""+Hostname+"\"}}\n")).Return(261, nil)
	mockFileForTestPlain := new(mock.File)
	mockFileForTestPlain.On("Write", []byte("plain.go:14: testChannel2:testChannel2 alert message "+uid+" []\n")).Return(93, nil)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				name: "testChannel",
				option: &Option{
					Output:       FILE,
					Format:       FormatJSON,
					File:         mockFileForTestJson,
					Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatusForLoggerNew(),
					MinimalLevel: Debug,
				},
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				name: "testChannel2",
				option: &Option{
					Output:       FILE,
					Format:       FormatPlain,
					File:         mockFileForTestPlain,
					Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatusForLoggerNew(),
					MinimalLevel: Debug,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.name, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.Alert(ctx, "message")
		})
	}
}
