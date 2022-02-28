package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/text/message"
)

var Log *zap.Logger

func init() {
	var err error
	config:=zap.NewProductionConfig()
	encoderConfig:=zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey="time stamp"
	encoderConfig.EncodeTime=zapcore.ISO8601TimeEncoder()
	encoderConfig.Stacktracekey=""
config.EncoderConfig=encoderConfig
log,err =config.Build(zap.AddCallerSkip(1))
	//Log, error = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}
func Info(Message string,fields ..zap.field){
	log.Info(message,fields)

}
func Debug(Message string,fields ..zap.field){
	log.Debug(message,fields)

}
func error(Message string,fields ..zap.field){
	log.error(message,fields)

}

