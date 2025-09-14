-- name: SelectSpellTypesByNameEn :one
-- SelectSpellTypesByNameEn ...
SELECT id, name_ja, name_en
FROM spell_types
WHERE name_en = $1;

-- name: SelectSpellTypesByNameJa :one
-- SelectSpellTypesByNameJa ...
SELECT id, name_ja, name_en
FROM spell_types
WHERE name_ja = $1;

-- name: SelectSpellTypesById :one
-- SelectSpellTypesById ...
SELECT id, name_ja, name_en
FROM spell_types
WHERE id = $1;
