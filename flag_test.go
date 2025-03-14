package logger

import (
	"log"
	"testing"
)

func createSettingInWhichAllFlagsHaveTheSameStatus(status bool) OptionFlags {
	return OptionFlags{
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
func createSettingInWhichAllButOneFlagsHaveStatusFalse(flag int) OptionFlags {
	o := createSettingInWhichAllFlagsHaveTheSameStatus(false)
	switch flag {
	case log.Ldate:
		o.Date = true
	case log.Ltime:
		o.Time = true
	case log.Lmicroseconds:
		o.Microseconds = true
	case log.Llongfile:
		o.LongFile = true
	case log.Lshortfile:
		o.ShortFile = true
	case log.LUTC:
		o.Utc = true
	case log.Lmsgprefix:
		o.MsgPrefix = true
	case log.LstdFlags:
		o.StdFlags = true
	}

	return o
}

func TestOptionFlags_IsDate(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.Ldate),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsDate(); got != tt.want {
				t.Errorf("IsDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionFlags_IsLongFile(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.Llongfile),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsLongFile(); got != tt.want {
				t.Errorf("IsLongFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionFlags_IsMicroseconds(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.Lmicroseconds),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsMicroseconds(); got != tt.want {
				t.Errorf("IsMicroseconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionFlags_IsMsgPrefix(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.Lmsgprefix),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsMsgPrefix(); got != tt.want {
				t.Errorf("IsMsgPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionFlags_IsShortFile(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.Lshortfile),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsShortFile(); got != tt.want {
				t.Errorf("IsShortFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionFlags_IsStdFlags(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.LstdFlags),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsStdFlags(); got != tt.want {
				t.Errorf("IsStdFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionFlags_IsTime(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.Ltime),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsTime(); got != tt.want {
				t.Errorf("IsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionFlags_IsUTC(t *testing.T) {
	tests := []struct {
		name   string
		fields OptionFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(true),
			want:   true,
		},
		{
			name:   "date off",
			fields: createSettingInWhichAllFlagsHaveTheSameStatus(false),
			want:   false,
		},
		{
			name:   "date one on",
			fields: createSettingInWhichAllButOneFlagsHaveStatusFalse(log.LUTC),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsUTC(); got != tt.want {
				t.Errorf("IsUTC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceFlags(t *testing.T) {
	ftr := createSettingInWhichAllFlagsHaveTheSameStatus(true)
	ffl := createSettingInWhichAllFlagsHaveTheSameStatus(false)
	tests := []struct {
		name   string
		fields InterfaceFlags
		want   bool
	}{
		{
			name:   "date on",
			fields: &ftr,
			want:   true,
		},
		{
			name:   "date off",
			fields: &ffl,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsDate(); got != tt.want {
				t.Errorf("IsDate() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsTime(); got != tt.want {
				t.Errorf("IsTime() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsMicroseconds(); got != tt.want {
				t.Errorf("IsMicroseconds() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsLongFile(); got != tt.want {
				t.Errorf("IsLongFile() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsShortFile(); got != tt.want {
				t.Errorf("IsShortFile() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsUTC(); got != tt.want {
				t.Errorf("IsUTC() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsMsgPrefix(); got != tt.want {
				t.Errorf("IsMsgPrefix() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsStdFlags(); got != tt.want {
				t.Errorf("IsStdFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
