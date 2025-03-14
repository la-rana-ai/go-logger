package logger

import (
	"testing"
)

func TestFormat_Num(t1 *testing.T) {
	tests := []struct {
		name   string
		fields Format
		want   int8
	}{
		{
			name:   "FormatPlain",
			fields: *FormatPlain,
			want:   0,
		},
		{
			name:   "FormatJSON",
			fields: *FormatJSON,
			want:   1,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.fields.Num(); got != tt.want {
				t1.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat_String(t1 *testing.T) {
	tests := []struct {
		name   string
		fields Format
		want   string
	}{
		{
			name:   "FormatPlain",
			fields: *FormatPlain,
			want:   "plain",
		},
		{
			name:   "FormatJSON",
			fields: *FormatJSON,
			want:   "json",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.fields.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFormat(t1 *testing.T) {
	tests := []struct {
		name    string
		fields  InterfaceFormat
		wantNum int8
		wantStr string
	}{
		{
			name:    "FormatPlain",
			fields:  FormatPlain,
			wantNum: 0,
			wantStr: "plain",
		},
		{
			name:    "FormatJSON",
			fields:  FormatJSON,
			wantNum: 1,
			wantStr: "json",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.fields.Num(); got != tt.wantNum {
				t1.Errorf("Num() = %v, want %v", got, tt.wantNum)
			}
			if got := tt.fields.String(); got != tt.wantStr {
				t1.Errorf("String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}
