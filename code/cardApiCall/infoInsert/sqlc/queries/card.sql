-- name: GetCard :one
SELECT * FROM cards
WHERE id = $1 LIMIT 1;

-- name: ListCards :many
SELECT * FROM cards
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: InsertCard :one
INSERT INTO cards (
  name_ja,
  name_en,
  card_text_ja,
  card_text_en,
  neuron_id,
  ocg_api_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateCard :one
UPDATE cards
SET
  name_ja = $2,
  name_en = $3,
  card_text_ja = $4,
  card_text_en = $5,
  neuron_id = $6,
  ocg_api_id = $7
WHERE id = $1
RETURNING *;

-- name: DeleteCard :exec
DELETE FROM cards
WHERE id = $1;

-- name: SelectByCardId :one
SELECT * FROM cards
WHERE id = $1;

-- name: SelectByCardNameEn :one
SELECT * FROM cards
WHERE name_en = $1;

-- name: SelectByCardNameJa :one
SELECT * FROM cards
WHERE name_ja = $1;

-- name: SelectCardPatternByCardID :one
WITH dummy AS (
    SELECT $1::bigint as id
)
SELECT
    dummy.id as card_id
    ,card.neuron_id
    ,card.ocg_api_id
    ,exists (select 1 from monsters where monsters.card_id = dummy.id) as is_monster
    ,exists (select 1 from spells where spells.card_id = dummy.id) as is_spell
    ,exists (select 1 from traps where traps.card_id = dummy.id) as is_trap
FROM
    dummy
LEFT JOIN
    cards AS card
    ON card.id = dummy.id;