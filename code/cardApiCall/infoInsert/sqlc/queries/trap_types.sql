-- name: SelectTrapTypesByNameEn :one
-- SelectTrapTypesByNameEn ...
SELECT id, name_ja, name_en
FROM trap_types
WHERE name_en = $1;

-- name: SelectTrapTypesByNameJa :one
-- SelectTrapTypesByNameJa ...
SELECT id, name_ja, name_en
FROM trap_types
WHERE name_ja = $1;

-- name: SelectTrapTypesById :one
-- SelectTrapTypesById ...
SELECT id, name_ja, name_en
FROM trap_types
WHERE id = $1;