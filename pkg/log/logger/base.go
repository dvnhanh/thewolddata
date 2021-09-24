package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type baselogger struct {
	logger zerolog.Logger
}

func (p *baselogger) initialize(writers ...io.Writer) {
	w := zerolog.MultiLevelWriter(writers...)
	event := zerolog.New(w).With().Timestamp()
	p.logger = event.Logger()
}

func (p *baselogger) writeLogger(logEvent *zerolog.Event, message string, data interface{}, err error) {
	if err != nil {
		logEvent.Interface("data", data).Err(err).Msg(message)
	}
	logEvent.Interface("data", data).Msg(message)
}
