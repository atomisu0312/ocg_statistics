package env

import (
	"os"
	"strconv"
)

// GetAsString 環境変数を文字列として取得します。存在しない場合はデフォルト値を返します。
func GetAsString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// GetAsInt 環境変数を整数として取得します。存在しない場合はデフォルト値を返します。
func GetAsInt(key string, defaultValue int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
