package logger

type InterfaceFormat interface {
	String() string
	Num() int8
}

type Format struct {
	int int8
	str string
}

/*
Num: logging type as a number
*/
func (t *Format) Num() int8 {
	return t.int
}

/*
String: logging type as a string
*/
func (t *Format) String() string {
	return t.str
}

var (
	/*
		FormatPlain: Output log to string
	*/
	FormatPlain = &Format{int: 0, str: "plain"}

	/*
		FormatPlain: Output log to json
	*/
	FormatJSON = &Format{int: 1, str: "json"}
)
