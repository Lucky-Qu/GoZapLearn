package logger

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	ZapLogger *zap.Logger
	LogLevel  logger.LogLevel
}

var GormLog GormLogger

func (g GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := g
	newLogger.LogLevel = level
	return &newLogger
}

func (g GormLogger) Info(ctx context.Context, msg string, i ...interface{}) {
	if g.LogLevel >= logger.Info {
		g.ZapLogger.Info(msg, zap.Any("data", i))
	}
}

func (g GormLogger) Warn(ctx context.Context, msg string, i ...interface{}) {
	if g.LogLevel >= logger.Warn {
		g.ZapLogger.Warn(msg, zap.Any("data", i))
	}
}

func (g GormLogger) Error(ctx context.Context, msg string, i ...interface{}) {
	if g.LogLevel >= logger.Error {
		g.ZapLogger.Error(msg, zap.Any("data", i))
	}
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		// 如果有错误，记录错误日志
		g.ZapLogger.Error("Error occurred",
			zap.Error(err),
			zap.String("sql", sql),
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
		)
	} else if elapsed > 200*time.Millisecond { //  200ms 作为慢查询的阈值
		g.ZapLogger.Warn("Slow query",
			zap.String("sql", sql),
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
		)
	} else {
		g.ZapLogger.Info("Query executed",
			zap.String("sql", sql),
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
		)
	}
}
