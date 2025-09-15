-- name: SelectAttributesByNameEn :one
-- SelectAttributesByNameEn ...
SELECT id, name_ja, name_en
FROM attributes
WHERE name_en = $1;

-- name: SelectAttributesByNameJa :one
-- SelectAttributesByNameJa ...
SELECT id, name_ja, name_en
FROM attributes
WHERE name_ja = $1;

-- name: SelectAttributesById :one
-- SelectAttributesById ...
SELECT id, name_ja, name_en
FROM attributes
WHERE id = $1;