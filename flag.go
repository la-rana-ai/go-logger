package logger

type InterfaceFlags interface {
	/*
	   IsDate: flag that sets whether to output the date in the local time zone: 2009/01/23
	*/
	IsDate() bool
	/*
	   IsTime: flag that sets whether to output the time in the local time zone: 01:23:23
	*/
	IsTime() bool
	/*
	   IsMicroseconds: flag that sets whether to output microsecond resolution: 01:23:23.123123.  assumes Time.
	*/
	IsMicroseconds() bool
	/*
	   IsLongFile: flag that sets whether to output full file name and line number: /a/b/c/d.go:23
	*/
	IsLongFile() bool
	/*
	   IsShortFile: flag that sets whether to output final file name element and line number: d.go:23. overrides LongFile
	*/
	IsShortFile() bool
	/*
	   IsUTC: A flag that determines whether UTC should be used instead of the local time zone when outputting a date or time.
	*/
	IsUTC() bool
	/*
	   IsMsgPrefix: flag that sets whether to output move the "prefix" from the beginning of the line to before the message
	*/
	IsMsgPrefix() bool
	/*
	   IsStdFlags: flag that sets whether to output initial values for the standard logger
	*/
	IsStdFlags() bool
}
type OptionFlags struct {
	Date         bool // the date in the local time zone: 2009/01/23
	Time         bool // the time in the local time zone: 01:23:23
	Microseconds bool // microsecond resolution: 01:23:23.123123.  assumes Time.
	LongFile     bool // full file name and line number: /a/b/c/d.go:23
	ShortFile    bool // final file name element and line number: d.go:23. overrides LLongFile
	Utc          bool // if Date or Time is set, use UTC rather than the local time zone
	MsgPrefix    bool // move the "prefix" from the beginning of the line to before the message
	StdFlags     bool // initial values for the standard logger
}

func (o *OptionFlags) IsDate() bool {
	return o.Date
}
func (o *OptionFlags) IsTime() bool {
	return o.Time
}
func (o *OptionFlags) IsMicroseconds() bool {
	return o.Microseconds
}
func (o *OptionFlags) IsLongFile() bool {
	return o.LongFile
}
func (o *OptionFlags) IsShortFile() bool {
	return o.ShortFile
}
func (o *OptionFlags) IsUTC() bool {
	return o.Utc
}
func (o *OptionFlags) IsMsgPrefix() bool {
	return o.MsgPrefix
}
func (o *OptionFlags) IsStdFlags() bool {
	return o.StdFlags
}
