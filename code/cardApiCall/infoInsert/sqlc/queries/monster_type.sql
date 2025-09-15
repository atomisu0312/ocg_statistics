-- name: SelectMonsterTypesByNameEn :one
-- SelectMonsterTypesByNameEn ...
SELECT id, name_ja, name_en
FROM monster_types
WHERE name_en = $1;

-- name: SelectMonsterTypesByNameJa :one
-- SelectMonsterTypesByNameJa ...
SELECT id, name_ja, name_en
FROM monster_types
WHERE name_ja = $1;

-- name: SelectMonsterTypesById :one
-- SelectMonsterTypesById ...
SELECT id, name_ja, name_en
FROM monster_types
WHERE id = $1;