-- name: FindSpellByCardID :one
-- FindSpellByCardID ...
SELECT card_id, spell_type_id
FROM spells
WHERE card_id = $1;

-- name: InsertSpell :one
-- InsertSpell ...
INSERT INTO spells (card_id, spell_type_id)
VALUES ($1, $2)
RETURNING *;

-- name: SelectFullSpellCardInfoByNeuronID :one
-- GetSpellCardByNeuronID ...
SELECT 
  cards.*
  , spell_types.name_ja as spell_type_name_ja
  , spell_types.name_en as spell_type_name_en
  FROM cards
join spells on cards.id = spells.card_id
join spell_types on spells.spell_type_id = spell_types.id
WHERE cards.neuron_id = $1;

-- name: SelectFullSpellCardInfoByOcgApiID :one
-- GetSpellCardByyOcgApiID ...
SELECT 
  cards.*
  , spell_types.name_ja as spell_type_name_ja
  , spell_types.name_en as spell_type_name_en
  FROM cards
join spells on cards.id = spells.card_id
join spell_types on spells.spell_type_id = spell_types.id
WHERE cards.ocg_api_id = $1;

-- name: SelectFullSpellCardInfoByCardID :one
-- GetSpellCardByyOcgApiID ...
SELECT 
  cards.*
  , spell_types.name_ja as spell_type_name_ja
  , spell_types.name_en as spell_type_name_en
  FROM cards
join spells on cards.id = spells.card_id
join spell_types on spells.spell_type_id = spell_types.id
WHERE cards.id = $1;
