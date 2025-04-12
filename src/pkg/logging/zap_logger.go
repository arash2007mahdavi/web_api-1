package logging

import (
	"github.com/arash2007mahdavi/web-api-1/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zaplogLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info": zapcore.InfoLevel,
	"warn": zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: l.cfg.Logger.FilePath,
		MaxSize: 1,
		MaxAge: 5,
		LocalTime: true,
		MaxBackups: 10,
		Compress: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	logger = logger.With("AppName", "MyApp", "LoggerName", "Zerolog")
	
	l.logger = logger
}


func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepaareLogKey(extra, cat, sub)

	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepaareLogKey(extra, cat, sub)

	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepaareLogKey(extra, cat, sub)

	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepaareLogKey(extra, cat, sub)

	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepaareLogKey(extra, cat, sub)

	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exists := zaplogLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func prepaareLogKey(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}

func mapToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)
	for k, v := range extra {
		params = append(params, string(k))
		params = append(params, v)
	}
	return params
}