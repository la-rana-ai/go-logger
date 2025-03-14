package logger

import (
	"github.com/google/uuid"
	"os"
)

const (
	DateSeparator         = "/"
	DateTimeSeparator     = "T"
	TimeSeparator         = ":"
	TimeNanoTimeSeparator = "."
	STDERR                = "stderr"
	STDOUT                = "stdout"
	DOCKERERR             = "d_err"
	DOCKEROUT             = "d_out"
	FILE                  = "f_port"
)

var (
	ProcessID, _ = uuid.NewUUID()
	Hostname, _  = os.Hostname()
)
