-- name: FindMonsterByCardID :one
-- FindMonsterByCardID ...
SELECT card_id, race_id, attribute_id, attack, defense, level, type_ids
FROM monsters
WHERE card_id = $1;

-- name: InsertMonster :one
-- InsertMonster ...
INSERT INTO monsters (card_id, race_id, attribute_id, attack, defense, level, type_ids)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: SelectFullMonsterCardInfoByOcgApiID :one
-- GetMonsterCardInfoByOcgApiID  ...
with target_card as (
    select
        *
    from
        cards
    where
        ocg_api_id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id;


-- name: SelectFullMonsterCardInfoByNeuronID :one
-- GetMonsterCardByNeuronID ...
with target_card as (
    select
        *
    from
        cards
    where
        neuron_id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id;


-- name: SelectFullMonsterCardInfoByCardID :one
-- GetMonsterCardByCardID ...
with target_card as (
    select
        *
    from
        cards
    where
        cards.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id;

-- name: InsertFusionMonster :one
-- InsertFusionMonster ...
INSERT INTO fusion_monsters (card_id)
VALUES ($1)
RETURNING *;

-- name: InsertSynchroMonster :one
-- InsertSynchroMonster ...
INSERT INTO synchro_monsters (card_id)
VALUES ($1)
RETURNING *;

-- name: InsertXyzMonster :one
-- InsertXyzMonster ...
INSERT INTO xyz_monsters (card_id)
VALUES ($1)
RETURNING *;

-- name: InsertLinkMonster :one
-- InsertLinkMonster ...
INSERT INTO link_monsters (card_id, link_marker)
VALUES ($1, $2)
RETURNING *;

-- name: InsertPendulumMonster :one
-- InsertPendulumMonster ...
INSERT INTO pendulum_monsters (card_id, scale, pendulum_text_ja, pendulum_text_en)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: InsertRitualMonster :one
-- InsertRitualMonster ...
INSERT INTO ritual_monsters (card_id)
VALUES ($1)
RETURNING *;

-- name: SelectFullFusionMonsterCardInfoByCardID :one
-- GetFusionMonsterCardByCardID ...
with target_card as (
    select
        *
    from
        cards
    where
        cards.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id
join
    fusion_monsters as fm on c.id = fm.card_id;

-- name: SelectFullSynchroMonsterCardInfoByCardID :one
-- GetSynchroMonsterCardByCardID ...
with target_card as (
    select
        *
    from
        cards
    where
        cards.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id
join
    synchro_monsters as sm on c.id = sm.card_id;

-- name: SelectFullXyzMonsterCardInfoByCardID :one
-- GetXyzMonsterCardByCardID ...
with target_card as (
    select
        *
    from
        cards
    where
        cards.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id
join
    xyz_monsters as xm on c.id = xm.card_id;

-- name: SelectFullLinkMonsterCardInfoByCardID :one
-- GetLinkMonsterCardByCardID ...
with target_card as (
    select
        *
    from
        cards
    where
        cards.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en,
    lm.link_marker
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id
join
    link_monsters as lm on c.id = lm.card_id;

-- name: SelectFullPendulumMonsterCardInfoByCardID :one
-- GetPendulumMonsterCardByCardID ...
with target_card as (
    select
        *
    from
        cards
    where
        cards.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en,
    pm.scale,
    pm.pendulum_text_ja,
    pm.pendulum_text_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id
join
    pendulum_monsters as pm on c.id = pm.card_id;

-- name: SelectFullRitualMonsterCardInfoByCardID :one
-- GetRitualMonsterCardByCardID ...
with target_card as (
    select
        *
    from
        cards
    where
        cards.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    c.*,
    m.attack,
    m.defense,
    m.level,
    ct.type_names_ja,
    ct.type_names_en,
    r.name_ja as race_name_ja,
    r.name_en as race_name_en,
    a.name_ja as attribute_name_ja,
    a.name_en as attribute_name_en
from
    target_card as c
join
    monsters as m on c.id = m.card_id
join
    races as r on m.race_id = r.id
join
    attributes as a on m.attribute_id = a.id
join
    card_types as ct on c.id = ct.card_id
join
    ritual_monsters as rm on c.id = rm.card_id;

-- name: SelectMonsterTypeLineByCardID :one
-- GetMonsterTypeLineByCardID ...
with target_card as (
    select
		c.id,
		c.neuron_id,
		c.ocg_api_id
    from
        cards c
    where
        c.id = $1
),
card_types as (
    select
        tt.id as card_id,
        array_agg(mt.name_ja)::varchar[] as type_names_ja,
        array_agg(mt.name_en)::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    cross join lateral
        unnest(m.type_ids) as t(type_id)
    join
        monster_types as mt on t.type_id = mt.id
    group by
        tt.id
)
select
    tc.id,
    tc.neuron_id,
    tc.ocg_api_id,
    'Normal' = ANY(ct.type_names_en)::boolean as is_normal,
    'Effect' = ANY(ct.type_names_en)::boolean as is_effect,
    'Toon' = ANY(ct.type_names_en)::boolean as is_toon,
    'Spirit' = ANY(ct.type_names_en)::boolean as is_spirit,
    'Union' = ANY(ct.type_names_en)::boolean as is_union,
    'Dual' = ANY(ct.type_names_en)::boolean as is_dual,
    'Tuner' = ANY(ct.type_names_en)::boolean as is_tuner,
    'Reverse' = ANY(ct.type_names_en)::boolean as is_reverse,
    rm.card_id is not null::boolean as is_ritual,
    xm.card_id is not null::boolean as is_xyz,
    sm.card_id is not null::boolean as is_synchro,
    fm.card_id is not null::boolean as is_fusion,
    lm.card_id is not null::boolean as is_link,
    pm.card_id is not null::boolean as is_pendulum
from 
    target_card tc
join card_types ct on tc.id = ct.card_id
left join ritual_monsters rm on rm.card_id  = tc.id
left join xyz_monsters xm on xm.card_id = tc.id
left join synchro_monsters sm on sm.card_id = tc.id
left join fusion_monsters fm on fm.card_id = tc.id
left join link_monsters lm on lm.card_id = tc.id
left join pendulum_monsters pm on pm.card_id = tc.id
;