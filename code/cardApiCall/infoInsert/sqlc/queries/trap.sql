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