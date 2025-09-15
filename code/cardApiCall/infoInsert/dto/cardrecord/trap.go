package cardrecord

import (
	"database/sql"

	"atomisu.com/ocg-statics/infoInsert/dto"
)

type TrapCardSelectResult struct {
	AbstractCardSelectResult
	TrapTypeNameJa string `db:"trap_type_name_ja" json:"trapTypeNameJa"`
	TrapTypeNameEn string `db:"trap_type_name_en" json:"trapTypeNameEn"`
}

type SelectFullTrapCardInfoRow struct {
	ID              int64          `db:"id" json:"id"`
	NeuronID        sql.NullInt64  `db:"neuron_id" json:"neuronId"`
	OcgApiID        sql.NullInt64  `db:"ocg_api_id" json:"ocgApiId"`
	NameJa          sql.NullString `db:"name_ja" json:"nameJa"`
	NameEn          sql.NullString `db:"name_en" json:"nameEn"`
	CardTextJa      sql.NullString `db:"card_text_ja" json:"cardTextJa"`
	CardTextEn      sql.NullString `db:"card_text_en" json:"cardTextEn"`
	Dataowner       sql.NullString `db:"dataowner" json:"dataowner"`
	RegistDate      sql.NullTime   `db:"regist_date" json:"registDate"`
	EnableStartDate sql.NullTime   `db:"enable_start_date" json:"enableStartDate"`
	EnableEndDate   sql.NullTime   `db:"enable_end_date" json:"enableEndDate"`
	Version         sql.NullInt64  `db:"version" json:"version"`
	TrapTypeNameJa  sql.NullString `db:"trap_type_name_ja" json:"trapTypeNameJa"`
	TrapTypeNameEn  sql.NullString `db:"trap_type_name_en" json:"trapTypeNameEn"`
}

func (t *TrapCardSelectResult) FromSelectFullTrapCardInfoRow(row SelectFullTrapCardInfoRow) *TrapCardSelectResult {
	return &TrapCardSelectResult{
		AbstractCardSelectResult: AbstractCardSelectResult{
			AbstractSelectResult: dto.AbstractSelectResult{
				ID:              row.ID,
				Dataowner:       row.Dataowner.String,
				RegistDate:      row.RegistDate.Time,
				EnableStartDate: row.EnableStartDate.Time,
				EnableEndDate:   row.EnableEndDate.Time,
				Version:         row.Version.Int64,
			},
			NeuronID:   row.NeuronID.Int64,
			OcgApiID:   row.OcgApiID.Int64,
			NameJa:     row.NameJa.String,
			NameEn:     row.NameEn.String,
			CardTextJa: row.CardTextJa.String,
			CardTextEn: row.CardTextEn.String,
		},
		TrapTypeNameJa: row.TrapTypeNameJa.String,
		TrapTypeNameEn: row.TrapTypeNameEn.String,
	}
}
