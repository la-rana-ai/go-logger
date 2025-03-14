package logger

import "testing"

func TestLevel_Num(t *testing.T) {
	tests := []struct {
		name   string
		fields Level
		want   uint8
	}{
		{
			name: "zero",
			fields: Level{
				int: 0,
				str: "zero",
			},
			want: 0,
		},
		{
			name: "one",
			fields: Level{
				int: 1,
				str: "one",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Num(); got != tt.want {
				t.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name   string
		fields Level
		want   string
	}{
		{
			name: "zero",
			fields: Level{
				int: 0,
				str: "zero",
			},
			want: "zero",
		},
		{
			name: "one",
			fields: Level{
				int: 1,
				str: "one",
			},
			want: "one",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceLevel(t *testing.T) {

	tests := []struct {
		name    string
		fields  InterfaceLevel
		wantNum uint8
		wantStr string
	}{
		{
			name:    "Debug",
			fields:  Debug,
			wantNum: 7,
			wantStr: "debug",
		},
		{
			name:    "Info",
			fields:  Info,
			wantNum: 6,
			wantStr: "info",
		},
		{
			name:    "Notice",
			fields:  Notice,
			wantNum: 5,
			wantStr: "notice",
		},
		{
			name:    "Warning",
			fields:  Warning,
			wantNum: 4,
			wantStr: "warning",
		},
		{
			name:    "Error",
			fields:  Error,
			wantNum: 3,
			wantStr: "error",
		},
		{
			name:    "Critical",
			fields:  Critical,
			wantNum: 2,
			wantStr: "critical",
		},
		{
			name:    "Alert",
			fields:  Alert,
			wantNum: 1,
			wantStr: "alert",
		},
		{
			name:    "Emergency",
			fields:  Emergency,
			wantNum: 0,
			wantStr: "emergency",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Num(); got != tt.wantNum {
				t.Errorf("Num() = %v, want %v", got, tt.wantNum)
			}
			if got := tt.fields.String(); got != tt.wantStr {
				t.Errorf("String() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}
