package logger

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Context struct {
	Args      any       `json:"args,omitempty"`
	RequestID any       `json:"x-request-id,omitempty"`
	ProcessID uuid.UUID `json:"process-id,omitempty"`
	UserAgent any       `json:"user-agent,omitempty"`
	Host      string    `json:"host,omitempty"`
}
type Datetime struct {
	Date     string `json:"date,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

type Log struct {
	Channel   string    `json:"channel,omitempty"`
	LevelName string    `json:"level-name,omitempty"`
	Level     uint8     `json:"level,omitempty"`
	File      string    `json:"file,omitempty"`
	Line      uint      `json:"line,omitempty"`
	Message   string    `json:"message,omitempty"`
	Context   Context   `json:"context,omitempty"`
	Datetime  *Datetime `json:"datetime,omitempty"`
}

/*
itoa: translate a number into a string with a specified minimum length
*/
func itoa(val int, min uint) string {
	str := strconv.Itoa(val)
	diff := int(min) - len(str)
	if diff > 0 {
		str = strings.Repeat("0", diff) + str
	}
	return str
}

/*
getFileAndLine: getting the file address and line number from where this log was called
*/
func getFileAndLine(flags int, pc int) (string, uint) {
	if flags&(log.Lshortfile|log.Llongfile) != 0 {
		_, file, line, ok := runtime.Caller(pc + 1)
		if !ok {
			file = "???"
			line = 0
		}
		if flags&log.Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		return file, uint(line)
	}
	return "", 0
}

/*
getDatetime: getting a human-readable timestamp
*/
func getDatetime(flags int) *Datetime {
	var (
		t  time.Time
		dt Datetime
	)
	if flags&(log.Ldate|log.Ltime|log.Lmicroseconds) != 0 {
		if flags&log.LUTC != 0 {
			t = t.UTC()
			dt.Timezone = "UTC"
		}
		if flags&log.Ldate != 0 {
			year, month, day := t.Date()
			dt.Date = itoa(year, 4) + DateSeparator + itoa(int(month), 2) + DateSeparator + itoa(day, 2)
			if flags&log.Ltime != 0 {
				dt.Date += DateTimeSeparator
			}
		}
		if flags&(log.Ltime|log.Lmicroseconds) != 0 {
			hour, minutes, sec := t.Clock()
			dt.Date += itoa(hour, 2) + TimeSeparator + itoa(minutes, 2) + TimeSeparator + itoa(sec, 2)
			if flags&log.Lmicroseconds != 0 {
				dt.Date += TimeNanoTimeSeparator + itoa(t.Nanosecond()/1e3, 6)
			}
		}
		return &dt
	}

	return nil
}

/*
traceJson: the main function of generating a log in json format
*/
func traceJson(logger *Logger, ctx context.Context, level InterfaceLevel, message string, context ...any) {
	if logger.minimalLevel != nil && logger.minimalLevel.Num() < level.Num() {
		return
	}
	l := Log{
		Channel:   logger.channel,
		LevelName: level.String(),
		Level:     level.Num(),
		Message:   message,
		Context: Context{
			Args:      context,
			RequestID: ctx.Value("x-request-id"),
			ProcessID: ProcessID,
			UserAgent: ctx.Value("user-agent"),
			Host:      Hostname,
		},
		Datetime: getDatetime(logger.flags),
	}

	l.File, l.Line = getFileAndLine(logger.flags, 2)

	ls, _ := json.Marshal(l)
	logger.log.Println(string(ls))
}
