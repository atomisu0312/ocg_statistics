package carddto

import (
	"database/sql"

	"atomisu.com/ocg-statics/infoInsert/dto"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
)

// StandardCard はUsecaseの関数の引数として機能する
type StandardCard struct {
	DescEn         string   `json:"descEn"`
	DescJa         string   `json:"descJa"`
	NameEn         string   `json:"nameEn"`
	NameJa         string   `json:"nameJa"`
	NeuronID       int64    `json:"neuronId"`
	TcgID          int64    `json:"tcgId"`
	Def            int64    `json:"def"`
	Atk            int64    `json:"atk"`
	Type           string   `json:"type"`
	Level          int64    `json:"level"`
	Race           string   `json:"race"`
	LinkMarkers    []string `json:"linkmarkers"`
	Attribute      string   `json:"attribute"`
	LinkVal        int64    `json:"linkval"`
	TypeLines      []string `json:"typeline"`
	CardType       string   `json:"cardType"`
	PendulumTextJa string   `json:"pendulumTextJa"`
	PendulumTextEn string   `json:"pendulumTextEn"`
}

func (s *StandardCard) ToInsertCardParamsExceptMonster() sqlc_gen.InsertCardParams {
	return sqlc_gen.InsertCardParams{
		NameEn:     sql.NullString{String: s.NameEn, Valid: true},
		NameJa:     sql.NullString{String: s.NameJa, Valid: true},
		CardTextEn: sql.NullString{String: s.DescEn, Valid: true},
		CardTextJa: sql.NullString{String: s.DescJa, Valid: true},
		NeuronID:   sql.NullInt64{Int64: s.NeuronID, Valid: true},
		OcgApiID:   sql.NullInt64{Int64: s.TcgID, Valid: true},
	}
}

type AbstractCardSelectResult struct {
	*dto.AbstractSelectResult
	NeuronID   int64  `db:"neuron_id" json:"neuronId"`
	OcgApiID   int64  `db:"ocg_api_id" json:"ocgApiId"`
	NameJa     string `db:"name_ja" json:"nameJa"`
	NameEn     string `db:"name_en" json:"nameEn"`
	CardTextJa string `db:"card_text_ja" json:"cardTextJa"`
	CardTextEn string `db:"card_text_en" json:"cardTextEn"`
}
