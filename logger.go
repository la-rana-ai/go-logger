package logger

import (
	"context"
	"log"
)

type Logger struct {
	minimalLevel InterfaceLevel
	flags        int
	log          *log.Logger
	channel      string
	trace        func(logger *Logger, ctx context.Context, level InterfaceLevel, message string, context ...any)
}

/*
Debug: (100) Detailed debug information.
*/
func (l *Logger) Debug(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Debug, message, context...)
}

/*
Info: (200) Interesting events. Examples: User logs in, SQL logs.
*/
func (l *Logger) Info(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Info, message, context...)
}

/*
Notice: (250) Normal but significant events.
*/
func (l *Logger) Notice(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Notice, message, context...)
}

/*
Warning: (300) Exceptional occurrences that are not errors.
Examples: Use of deprecated APIs, poor use of an API, undesirable things that are not necessarily wrong.
*/
func (l *Logger) Warning(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Warning, message, context...)
}

/*
Error: (400) Runtime errors that do not require immediate action but should typically be logged and monitored.
*/
func (l *Logger) Error(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Error, message, context...)
}

/*
Critical: (500) conditions. Example: Application component unavailable, unexpected exception.
*/
func (l *Logger) Critical(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Critical, message, context...)
}

/*
Alert: (550) Action must be taken immediately. Example: Entire website down, database unavailable, etc.
This should trigger the SMS alerts and wake you up.
*/
func (l *Logger) Alert(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Alert, message, context...)
}

/*
Emergency: (600) system is unusable.
*/
func (l *Logger) Emergency(ctx context.Context, message string, context ...any) {
	l.trace(l, ctx, Emergency, message, context...)
}
