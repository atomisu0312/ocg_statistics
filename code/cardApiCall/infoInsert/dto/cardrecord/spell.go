package cardrecord

import (
	"database/sql"

	"atomisu.com/ocg-statics/infoInsert/dto"
)

type SpellCardSelectResult struct {
	AbstractCardSelectResult
	SpellTypeNameJa string `db:"spell_type_name_ja" json:"spellTypeNameJa"`
	SpellTypeNameEn string `db:"spell_type_name_en" json:"spellTypeNameEn"`
}

type SelectFullSpellCardInfoRow struct {
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
	SpellTypeNameJa sql.NullString `db:"spell_type_name_ja" json:"spellTypeNameJa"`
	SpellTypeNameEn sql.NullString `db:"spell_type_name_en" json:"spellTypeNameEn"`
}

func (s *SpellCardSelectResult) FromSelectFullSpellCardInfoRow(row SelectFullSpellCardInfoRow) *SpellCardSelectResult {
	return &SpellCardSelectResult{
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
		SpellTypeNameJa: row.SpellTypeNameJa.String,
		SpellTypeNameEn: row.SpellTypeNameEn.String,
	}
}
