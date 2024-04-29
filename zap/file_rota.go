package base_zap

import (
	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

func (r *fileRotatelogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	//fileWriter := NewCutter("./", level, WithCutterFormat("2006-01-02"))
	//if global.GVA_CONFIG.Zap.LogInConsole {
	//	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	//}
	//return zapcore.AddSync(fileWriter)
	return nil
}
