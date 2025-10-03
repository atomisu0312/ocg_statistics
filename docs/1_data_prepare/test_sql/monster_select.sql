with target_card as (
    select
        *
    from
        cards
),
card_types as (
    select
        tt.id as card_id,
        COALESCE(array_agg(mt.name_ja) FILTER (WHERE t.type_id IS NOT NULL), '{}')::varchar[] as type_names_ja,
        COALESCE(array_agg(mt.name_en) FILTER (WHERE t.type_id IS NOT NULL), '{}')::varchar[] as type_names_en
    from
        target_card as tt
    join
    monsters as m on m.card_id = tt.id
    left join lateral
        unnest(m.type_ids) as t(type_id) on true
    left join
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
left join
    card_types as ct on c.id = ct.card_id;