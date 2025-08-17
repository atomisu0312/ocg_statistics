-- name: FindTrapByCardID :one
-- FindTrapByCardID ...
SELECT card_id, trap_type_id
FROM traps
WHERE card_id = $1;