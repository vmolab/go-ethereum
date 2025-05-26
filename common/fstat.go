package common

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

var CallLogger io.Writer

const DoCallLog bool = true

func init() {
	CallLogger = &lumberjack.Logger{
		Filename: "log/evm_calls.log",
		MaxSize:  512, // megabytes
		Compress: true,
	}
}
