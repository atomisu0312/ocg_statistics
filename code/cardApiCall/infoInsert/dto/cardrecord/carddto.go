package cardrecord

import (
	"atomisu.com/ocg-statics/infoInsert/dto"
	"atomisu.com/ocg-statics/infoInsert/sqlc_gen"
	"atomisu.com/ocg-statics/infoInsert/util"
)

// StandardCard はUsecaseの関数の引数として機能する
type StandardCard struct {
	DescEn         string   `json:"descEn"`
	DescJa         string   `json:"descJa"`
	NameEn         string   `json:"nameEn"`
	NameJa         string   `json:"nameJa"`
	NeuronID       int64    `json:"neuronId"`
	TcgID          int64    `json:"tcgId"`
	Def            int32    `json:"def"`
	Atk            int32    `json:"atk"`
	Type           string   `json:"type"`
	Level          int32    `json:"level"`
	Race           string   `json:"race"`
	LinkMarkers    []string `json:"linkMarkers"`
	Attribute      string   `json:"attribute"`
	LinkVal        int32    `json:"linkVal"`
	TypeLines      []string `json:"typeLines"`
	CardType       string   `json:"cardType"`
	PendulumScale  int32    `json:"pendulumScale"`
	PendulumTextJa string   `json:"pendulumTextJa"`
	PendulumTextEn string   `json:"pendulumTextEn"`
}

func (s *StandardCard) ToInsertCardParamsExceptMonster() sqlc_gen.InsertCardParams {
	return sqlc_gen.InsertCardParams{
		NameEn:     util.ParseAsSqlNullString(s.NameEn),
		NameJa:     util.ParseAsSqlNullString(s.NameJa),
		CardTextEn: util.ParseAsSqlNullString(s.DescEn),
		CardTextJa: util.ParseAsSqlNullString(s.DescJa),
		NeuronID:   util.ParseAsSqlNullInt64WithTreatZeroAsNull(s.NeuronID),
		OcgApiID:   util.ParseAsSqlNullInt64WithTreatZeroAsNull(s.TcgID),
	}
}

type AbstractCardSelectResult struct {
	dto.AbstractSelectResult
	NeuronID   int64  `db:"neuron_id" json:"neuronId"`
	OcgApiID   int64  `db:"ocg_api_id" json:"ocgApiId"`
	NameJa     string `db:"name_ja" json:"nameJa"`
	NameEn     string `db:"name_en" json:"nameEn"`
	CardTextJa string `db:"card_text_ja" json:"cardTextJa"`
	CardTextEn string `db:"card_text_en" json:"cardTextEn"`
}
