package logger

import (
	"github.com/la-rana-ai/go-logger/mock"
	"io"
	"reflect"
	"testing"
)

func createPointerSettingInWhichAllFlagsHaveTheSameStatus(status bool) *OptionFlags {
	return &OptionFlags{
		Date:         status,
		Time:         status,
		Microseconds: status,
		LongFile:     status,
		ShortFile:    status,
		Utc:          status,
		MsgPrefix:    status,
		StdFlags:     status,
	}
}

func TestOption_GetFile(t *testing.T) {
	mockFile := new(mock.File)

	tests := []struct {
		name   string
		fields Option
		want   io.Writer
	}{
		{
			name: "stderr",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
				File:         mockFile,
			},
			want: nil,
		},
		{
			name: "stderr and file descriptor",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: nil,
		},
		{
			name: "stdout",
			fields: Option{
				Output:       STDOUT,
				Format:       FormatJSON,
				File:         mockFile,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: nil,
		},
		{
			name: "stdout",
			fields: Option{
				Output:       STDOUT,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: nil,
		},
		{
			name: "d_err",
			fields: Option{
				Output:       DOCKERERR,
				Format:       FormatJSON,
				File:         mockFile,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: nil,
		},
		{
			name: "d_out",
			fields: Option{
				Output:       DOCKEROUT,
				Format:       FormatJSON,
				File:         mockFile,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: nil,
		},
		{
			name: "f_port",
			fields: Option{
				Output:       FILE,
				Format:       FormatJSON,
				File:         mockFile,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: mockFile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GetFlags(t *testing.T) {
	ff := createPointerSettingInWhichAllFlagsHaveTheSameStatus(false)
	ft := createPointerSettingInWhichAllFlagsHaveTheSameStatus(true)
	f := createPointerSettingInWhichAllFlagsHaveTheSameStatus(true)
	f.LongFile = false
	tests := []struct {
		name   string
		fields Option
		want   InterfaceFlags
	}{
		{
			name: "all false",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        ff,
				MinimalLevel: Debug,
			},
			want: ff,
		},
		{
			name: "all true",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        ft,
				MinimalLevel: Debug,
			},
			want: ft,
		},
		{
			name: "LongFile false, other true",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        f,
				MinimalLevel: Debug,
			},
			want: f,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetFlags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GetFormat(t *testing.T) {
	tests := []struct {
		name   string
		fields Option
		want   InterfaceFormat
	}{
		{
			name: "FormatJSON",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: FormatJSON,
		},
		{
			name: "FormatPlain",
			fields: Option{
				Output:       STDERR,
				Format:       FormatPlain,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: FormatPlain,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetFormat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GetMinimalLevel(t *testing.T) {
	tests := []struct {
		name   string
		fields Option
		want   InterfaceLevel
	}{
		{
			name: "Debug",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: Debug,
		},
		{
			name: "Info",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Info,
			},
			want: Info,
		},
		{
			name: "Notice",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Notice,
			},
			want: Notice,
		},
		{
			name: "Warning",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Warning,
			},
			want: Warning,
		},
		{
			name: "Error",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Error,
			},
			want: Error,
		},
		{
			name: "Critical",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Critical,
			},
			want: Critical,
		},
		{
			name: "Alert",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Alert,
			},
			want: Alert,
		},
		{
			name: "Emergency",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Emergency,
			},
			want: Emergency,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetMinimalLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinimalLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GetOutput(t *testing.T) {
	tests := []struct {
		name   string
		fields Option
		want   string
	}{
		{
			name: "stderr",
			fields: Option{
				Output:       STDERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: STDERR,
		},
		{
			name: "stdout",
			fields: Option{
				Output:       STDOUT,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: STDOUT,
		},
		{
			name: "d_err",
			fields: Option{
				Output:       DOCKERERR,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: DOCKERERR,
		},
		{
			name: "d_out",
			fields: Option{
				Output:       DOCKEROUT,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: DOCKEROUT,
		},
		{
			name: "f_port",
			fields: Option{
				Output:       FILE,
				Format:       FormatJSON,
				Flags:        createPointerSettingInWhichAllFlagsHaveTheSameStatus(false),
				MinimalLevel: Debug,
			},
			want: FILE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetOutput(); got != tt.want {
				t.Errorf("GetOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
