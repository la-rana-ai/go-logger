package logger

import (
	"context"
)

/*
tracePlain: the main function of generating a log in plain format
*/
func tracePlain(logger *Logger, ctx context.Context, level InterfaceLevel, message string, context ...any) {
	if logger.minimalLevel != nil && logger.minimalLevel.Num() < level.Num() {
		return
	}
	logger.log.Println(logger.channel, level.String(), message, ctx.Value(`x-request-id`), context)
}
