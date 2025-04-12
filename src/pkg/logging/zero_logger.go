package logging

import (
	"os"

	"github.com/arash2007mahdavi/web-api-1/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zerologLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info": zerolog.InfoLevel,
	"warn": zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

type zeroLogger struct {
	cfg *config.Config
	logger *zerolog.Logger
}

func NewZeroLogger(cfg *config.Config) *zeroLogger {
	l := &zeroLogger{}
	l.cfg = cfg
	l.Init()
	return l
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, ok := zerologLevelMap[l.cfg.Logger.Level]
	if ok {
		return level
	}
	return zerolog.DebugLevel
}

func (l *zeroLogger) Init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("could not open log file")
	}

	var logger = zerolog.New(file).
	With().
	Timestamp().
	Str("AppName", "MyApp").
	Str("LoggerName", "Zerolog").
	Logger()

	zerolog.SetGlobalLevel(l.getLogLevel())

	l.logger = &logger
}

func logParamsTozeroParams(keys map[ExtraKey]interface{}) map[string]interface{} {
	params := make(map[string]interface{})

	for k, v := range keys {
		params[string(k)] = v
	}

	return params
}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := logParamsTozeroParams(extra)

	l.logger.Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := logParamsTozeroParams(extra)

	l.logger.Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := logParamsTozeroParams(extra)

	l.logger.Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := logParamsTozeroParams(extra)

	l.logger.Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := logParamsTozeroParams(extra)

	l.logger.Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(params).
		Msg(msg)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}