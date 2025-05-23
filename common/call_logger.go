package common

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

var CallLogger io.Writer

func init() {
	CallLogger = &lumberjack.Logger{
		Filename: "evm_calls.log",
		MaxSize:  8, // megabytes
		Compress: false,
	}
}
