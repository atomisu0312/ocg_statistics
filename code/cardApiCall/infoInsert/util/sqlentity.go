package util

import "database/sql"

func ParseAsSqlNullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: v, Valid: true}
}

func ParseAsSqlNullInt64WithTreatZeroAsNull(v int64) sql.NullInt64 {
	// 0 を有効な値として扱うかは仕様次第。0 を NULL としたいなら `v == 0` を NULL 判定にする。
	if v < 1 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: v, Valid: true}
}

func ParseAsSqlNullInt32WithTreatZeroAsNull(v int32) sql.NullInt32 {
	// 0 を有効な値として扱うかは仕様次第。0 を NULL としたいなら `v == 0` を NULL 判定にする。
	if v < 1 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: v, Valid: true}
}
