package logger

import (
	"io"
	"os"

	"github.com/dvnhanh/thewolddata/pkg/log/logger/entity"
)

func NewConsoleLogger(filters ...io.Writer) Logger {
	p := new(consoleLog)
	filters = append(filters, os.Stdout)
	p.initialize(filters...)
	return p
}

type consoleLog struct {
	baselogger
}

func (p *consoleLog) Debug(log *entity.LogItem) {
	p.writeLogger(p.logger.Debug(), log.Message, log.Data, log.Error)
}

func (p *consoleLog) Info(log *entity.LogItem) {
	p.writeLogger(p.logger.Info(), log.Message, log.Data, log.Error)
}

func (p *consoleLog) Warn(log *entity.LogItem) {
	p.writeLogger(p.logger.Warn(), log.Message, log.Data, log.Error)
}

func (p *consoleLog) Error(log *entity.LogItem) {
	p.writeLogger(p.logger.Error(), log.Message, log.Data, log.Error)
}

func (p *consoleLog) Fatal(log *entity.LogItem) {
	p.writeLogger(p.logger.Fatal(), log.Message, log.Data, log.Error)
}
