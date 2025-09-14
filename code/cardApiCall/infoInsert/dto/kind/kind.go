package kind

import "database/sql"

type Kind struct {
	ID     int32
	NameJa string
	NameEn string
}

type SelectFullKindInfoRow struct {
	ID     int32          `db:"id" json:"id"`
	NameJa sql.NullString `db:"name_ja" json:"nameJa"`
	NameEn sql.NullString `db:"name_en" json:"nameEn"`
}

func (k *Kind) FromSelectFullKindInfoRow(row SelectFullKindInfoRow) Kind {
	return Kind{
		ID:     row.ID,
		NameJa: row.NameJa.String,
		NameEn: row.NameEn.String,
	}
}
