package development

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Log struct{}

const (
	_FEATURE = "feature"
	_OP      = "op"
)

func (c Log) New(feature string, oper string, flags ...string) zerolog.Logger {
	logger := log.Logger.With().Str(_FEATURE, feature).Str(_OP, oper)
	for i := range flags {
		logger = logger.Bool(flags[i], true)
	}

	return logger.Logger()
}
