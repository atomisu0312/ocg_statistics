-- name: SelectRacesByNameEn :one
-- SelectRacesByNameEn ...
SELECT id, name_ja, name_en
FROM races
WHERE name_en = $1;

-- name: SelectRacesByNameJa :one
-- SelectRacesByNameJa ...
SELECT id, name_ja, name_en
FROM races
WHERE name_ja = $1;

-- name: SelectRacesById :one
-- SelectRacesById ...
SELECT id, name_ja, name_en
FROM races
WHERE id = $1;