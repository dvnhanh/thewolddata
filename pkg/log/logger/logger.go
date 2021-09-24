package logger

import "github.com/dvnhanh/thewolddata/pkg/log/logger/entity"

type Logger interface {
	Debug(Log *entity.LogItem)
	Info(Log *entity.LogItem)
	Warn(Log *entity.LogItem)
	Error(Log *entity.LogItem)
	Fatal(Log *entity.LogItem)
}
