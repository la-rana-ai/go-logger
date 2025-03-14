package logger

import (
	"io"
)

type InterfaceOption interface {
	GetOutput() string
	GetFormat() InterfaceFormat
	GetFile() io.Writer
	GetMinimalLevel() InterfaceLevel
	GetFlags() InterfaceFlags
}
type Option struct {
	Output       string
	Format       InterfaceFormat
	Flags        InterfaceFlags
	MinimalLevel InterfaceLevel
	File         io.Writer
}

func (o *Option) GetOutput() string {
	return o.Output
}
func (o *Option) GetFile() io.Writer {
	if o.Output != FILE {
		return nil
	}
	return o.File
}
func (o *Option) GetFlags() InterfaceFlags {
	return o.Flags
}
func (o *Option) GetFormat() InterfaceFormat {
	return o.Format
}
func (o *Option) GetMinimalLevel() InterfaceLevel {
	return o.MinimalLevel
}
