package logger

type InterfaceLevel interface {
	String() string
	Num() uint8
}
type Level struct {
	int uint8
	str string
}

/*
Num: logging level as a number
*/
func (l *Level) Num() uint8 {
	return l.int
}

/*
String: logging level as a string
*/
func (l *Level) String() string {
	return l.str
}

var (
	/*
	   Debug: (100) Detailed debug information.
	*/
	Debug = &Level{int: 7, str: "debug"}

	/*
		Info: (200) Interesting events. Examples: User logs in, SQL logs.
	*/
	Info = &Level{int: 6, str: "info"}

	/*
		Notice: (250) Normal but significant events.
	*/
	Notice = &Level{int: 5, str: "notice"}

	/*
		Warning: (300) Exceptional occurrences that are not errors.
		Examples: Use of deprecated APIs, poor use of an API, undesirable things that are not necessarily wrong.
	*/
	Warning = &Level{int: 4, str: "warning"}

	/*
		Error: (400) Runtime errors that do not require immediate action but should typically be logged and monitored.
	*/
	Error = &Level{int: 3, str: "error"}

	/*
		Critical: (500) conditions. Example: Application component unavailable, unexpected exception.
	*/
	Critical = &Level{int: 2, str: "critical"}

	/*
		Alert: (550) Action must be taken immediately. Example: Entire website down, database unavailable, etc.
		This should trigger the SMS alerts and wake you up.
	*/
	Alert = &Level{int: 1, str: "alert"}

	/*
		Emergency: (600) system is unusable.
	*/
	Emergency = &Level{int: 0, str: "emergency"}
)
