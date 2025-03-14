package logger

import (
	"errors"
	"io"
	"log"
	"os"
)

var files = make(map[string]io.Writer)

func getFile(filePath string) (io.Writer, error) {
	if file, ok := files[filePath]; ok {
		return file, nil
	}
	if out, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		return nil, err
	} else {
		files[filePath] = out
		return out, nil
	}
}

func getFlag(fl InterfaceFlags) int {
	var (
		flags int
	)
	if fl.IsDate() {
		flags |= log.Ldate
	}
	if fl.IsTime() {
		flags |= log.Ltime
	}
	if fl.IsMicroseconds() {
		flags |= log.Lmicroseconds
	}
	if fl.IsLongFile() {
		flags |= log.Llongfile
	}
	if fl.IsShortFile() {
		flags |= log.Lshortfile
	}
	if fl.IsUTC() {
		flags |= log.LUTC
	}
	if fl.IsMsgPrefix() {
		flags |= log.Lmsgprefix
	}
	if fl.IsStdFlags() {
		flags |= log.LstdFlags
	}

	return flags
}

func getOutput(option string, file io.Writer) (io.Writer, error) {
	var (
		out io.Writer
		err error
	)
	if option != "" {
		switch option {
		case STDOUT:
			out = os.Stdout
		case STDERR:
			out = os.Stderr
		case DOCKEROUT:
			if out, err = getFile("/proc/1/fd/1"); err != nil {
				return nil, err
			}
		case DOCKERERR:
			if out, err = getFile("/proc/1/fd/2"); err != nil {
				return nil, err
			}
		case FILE:
			if file == nil {
				return nil, errors.New("file is nil")
			}
			out = file
		default:
			if out, err = getFile(option); err != nil {
				return nil, err
			}
		}
	} else {
		out = os.Stdout
	}

	return out, nil
}
