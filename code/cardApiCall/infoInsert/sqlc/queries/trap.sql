-- name: FindTrapByCardID :one
-- FindTrapByCardID ...
SELECT card_id, trap_type_id
FROM traps
WHERE card_id = $1;

-- name: InsertTrap :one
-- InsertTrap ...
INSERT INTO traps (card_id, trap_type_id)
VALUES ($1, $2)
RETURNING *;

-- name: SelectFullTrapCardInfoByNeuronID :one
-- GetTrapCardByNeuronID ...
SELECT 
  cards.*
  , trap_types.name_ja as trap_type_name_ja
  , trap_types.name_en as trap_type_name_en
  FROM cards
join traps on cards.id = traps.card_id
join trap_types on traps.trap_type_id = trap_types.id
WHERE cards.neuron_id = $1;

-- name: SelectFullTrapCardInfoByOcgApiID :one
-- GetTrapCardByyOcgApiID ...
SELECT 
  cards.*
  , trap_types.name_ja as trap_type_name_ja
  , trap_types.name_en as trap_type_name_en
  FROM cards
join traps on cards.id = traps.card_id
join trap_types on traps.trap_type_id = trap_types.id
WHERE cards.ocg_api_id = $1;

-- name: SelectFullTrapCardInfoByCardID :one
-- GetTrapCardByyOcgApiID ...
SELECT 
  cards.*
  , trap_types.name_ja as trap_type_name_ja
  , trap_types.name_en as trap_type_name_en
  FROM cards
join traps on cards.id = traps.card_id
join trap_types on traps.trap_type_id = trap_types.id
WHERE cards.id = $1;