package cardrecord

import (
	"database/sql"

	"atomisu.com/ocg-statics/infoInsert/dto"
)

type MonsterCardSelectResult struct {
	AbstractCardSelectResult
	Attack          int32    `db:"attack" json:"attack"`
	Defense         int32    `db:"defense" json:"defense"`
	Level           int32    `db:"level" json:"level"`
	TypeNamesJa     []string `db:"type_names_ja" json:"typeNamesJa"`
	TypeNamesEn     []string `db:"type_names_en" json:"typeNamesEn"`
	RaceNameJa      string   `db:"race_name_ja" json:"raceNameJa"`
	RaceNameEn      string   `db:"race_name_en" json:"raceNameEn"`
	AttributeNameJa string   `db:"attribute_name_ja" json:"attributeNameJa"`
	AttributeNameEn string   `db:"attribute_name_en" json:"attributeNameEn"`
}

type FusionMonsterSelectResult struct {
	MonsterCardSelectResult
}

type SynchroMonsterSelectResult struct {
	MonsterCardSelectResult
}

type XyzMonsterSelectResult struct {
	MonsterCardSelectResult
}

type LinkMonsterSelectResult struct {
	MonsterCardSelectResult
	LinkMarker sql.NullInt32 `db:"link_marker" json:"linkMarker"`
}

type PendulumMonsterSelectResult struct {
	MonsterCardSelectResult
	Scale          sql.NullInt32  `db:"scale" json:"scale"`
	PendulumTextJa sql.NullString `db:"pendulum_text_ja" json:"pendulumTextJa"`
	PendulumTextEn sql.NullString `db:"pendulum_text_en" json:"pendulumTextEn"`
}

type RitualMonsterSelectResult struct {
	MonsterCardSelectResult
}

type SelectFullMonsterCardInfoRow struct {
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
	Attack          sql.NullInt32  `db:"attack" json:"attack"`
	Defense         sql.NullInt32  `db:"defense" json:"defense"`
	Level           sql.NullInt32  `db:"level" json:"level"`
	TypeNamesJa     []string       `db:"type_names_ja" json:"typeNamesJa"`
	TypeNamesEn     []string       `db:"type_names_en" json:"typeNamesEn"`
	RaceNameJa      sql.NullString `db:"race_name_ja" json:"raceNameJa"`
	RaceNameEn      sql.NullString `db:"race_name_en" json:"raceNameEn"`
	AttributeNameJa sql.NullString `db:"attribute_name_ja" json:"attributeNameJa"`
	AttributeNameEn sql.NullString `db:"attribute_name_en" json:"attributeNameEn"`
}

type SelectFullFusionMonsterCardInfoRow struct {
	SelectFullMonsterCardInfoRow
}

type SelectFullSynchroMonsterCardInfoRow struct {
	SelectFullMonsterCardInfoRow
}

type SelectFullXyzMonsterCardInfoRow struct {
	SelectFullMonsterCardInfoRow
}

type SelectFullLinkMonsterCardInfoRow struct {
	SelectFullMonsterCardInfoRow
	LinkMarker sql.NullInt32 `db:"link_marker" json:"linkMarker"`
}

type SelectFullPendulumMonsterCardInfoRow struct {
	SelectFullMonsterCardInfoRow
	Scale          sql.NullInt32  `db:"scale" json:"scale"`
	PendulumTextJa sql.NullString `db:"pendulum_text_ja" json:"pendulumTextJa"`
	PendulumTextEn sql.NullString `db:"pendulum_text_en" json:"pendulumTextEn"`
}

type SelectFullRitualMonsterCardInfoRow struct {
	SelectFullMonsterCardInfoRow
}

func (m *MonsterCardSelectResult) FromSelectFullMonsterCardInfoRow(row SelectFullMonsterCardInfoRow) *MonsterCardSelectResult {
	return &MonsterCardSelectResult{
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
		Attack:          row.Attack.Int32,
		Defense:         row.Defense.Int32,
		Level:           row.Level.Int32,
		TypeNamesJa:     row.TypeNamesJa,
		TypeNamesEn:     row.TypeNamesEn,
		RaceNameJa:      row.RaceNameJa.String,
		RaceNameEn:      row.RaceNameEn.String,
		AttributeNameJa: row.AttributeNameJa.String,
		AttributeNameEn: row.AttributeNameEn.String,
	}
}

func (f *FusionMonsterSelectResult) FromSelectFullFusionMonsterCardInfoRow(row SelectFullFusionMonsterCardInfoRow) *FusionMonsterSelectResult {
	base := (&MonsterCardSelectResult{}).FromSelectFullMonsterCardInfoRow(row.SelectFullMonsterCardInfoRow)
	return &FusionMonsterSelectResult{
		MonsterCardSelectResult: *base,
	}
}

func (s *SynchroMonsterSelectResult) FromSelectFullSynchroMonsterCardInfoRow(row SelectFullSynchroMonsterCardInfoRow) *SynchroMonsterSelectResult {
	base := (&MonsterCardSelectResult{}).FromSelectFullMonsterCardInfoRow(row.SelectFullMonsterCardInfoRow)
	return &SynchroMonsterSelectResult{
		MonsterCardSelectResult: *base,
	}
}

func (x *XyzMonsterSelectResult) FromSelectFullXyzMonsterCardInfoRow(row SelectFullXyzMonsterCardInfoRow) *XyzMonsterSelectResult {
	base := (&MonsterCardSelectResult{}).FromSelectFullMonsterCardInfoRow(row.SelectFullMonsterCardInfoRow)
	return &XyzMonsterSelectResult{
		MonsterCardSelectResult: *base,
	}
}

func (l *LinkMonsterSelectResult) FromSelectFullLinkMonsterCardInfoRow(row SelectFullLinkMonsterCardInfoRow) *LinkMonsterSelectResult {
	base := (&MonsterCardSelectResult{}).FromSelectFullMonsterCardInfoRow(row.SelectFullMonsterCardInfoRow)
	return &LinkMonsterSelectResult{
		MonsterCardSelectResult: *base,
		LinkMarker:              row.LinkMarker,
	}
}

func (p *PendulumMonsterSelectResult) FromSelectFullPendulumMonsterCardInfoRow(row SelectFullPendulumMonsterCardInfoRow) *PendulumMonsterSelectResult {
	base := (&MonsterCardSelectResult{}).FromSelectFullMonsterCardInfoRow(row.SelectFullMonsterCardInfoRow)
	return &PendulumMonsterSelectResult{
		MonsterCardSelectResult: *base,
		Scale:                   row.Scale,
		PendulumTextJa:          row.PendulumTextJa,
		PendulumTextEn:          row.PendulumTextEn,
	}
}

func (r *RitualMonsterSelectResult) FromSelectFullRitualMonsterCardInfoRow(row SelectFullRitualMonsterCardInfoRow) *RitualMonsterSelectResult {
	base := (&MonsterCardSelectResult{}).FromSelectFullMonsterCardInfoRow(row.SelectFullMonsterCardInfoRow)
	return &RitualMonsterSelectResult{
		MonsterCardSelectResult: *base,
	}
}
