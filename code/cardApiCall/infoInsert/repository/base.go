package repository

import (
	"time"

	"go.uber.org/zap"
)

// Repository は、リポジトリの基本インターフェースです。
type Repository interface {
	// 共通のメソッドがあれば追加
}

// repository は、リポジトリの基本構造体です。
type repository struct {
	logger *zap.Logger
}

// NewRepository は、リポジトリの新しいインスタンスを作成します。
func NewRepository[T Repository](constructor func(*repository) T) T {
	logger, _ := zap.NewDevelopment()
	r := &repository{logger: logger}
	return constructor(r)
}

// logDBOperation は、データベース操作のログを記録します。
func (r *repository) logDBOperation(operation string, start time.Time, fields ...zap.Field) {
	duration := time.Since(start)
	fields = append(fields, zap.Duration("duration", duration))
	r.logger.Info("Database operation completed", append([]zap.Field{zap.String("operation", operation)}, fields...)...)
}

// logDBError は、データベース操作のエラーを記録します。
func (r *repository) logDBError(operation string, err error, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))
	r.logger.Error("Database operation failed", append([]zap.Field{zap.String("operation", operation)}, fields...)...)
}

// logDBResult は、データベース操作の結果を記録します。
func (r *repository) logDBResult(operation string, result interface{}, fields ...zap.Field) {
	fields = append(fields, zap.Any("result", result))
	r.logger.Debug("Database operation result", append([]zap.Field{zap.String("operation", operation)}, fields...)...)
}
