BEGIN;
truncate table cards restart identity cascade;

-- 既存のカードデータ
insert into
cards(neuron_id, ocg_api_id, name_ja, name_en, card_text_ja, card_text_en)
values
(11, 1, 'サンプルただのモンスター','Sample Ordinal', 'テキスト1', 'text11' ),
(12, 2, 'サンプル通常魔法','Spell12', 'テキスト2', 'text12' ),
(13, 3, 'サンプル通常罠','Trap13', 'テキスト3', 'text13' ),
(14, 4, 'サンプル速攻魔法','Trap14', 'テキスト4', 'text14' ),
(15, 5, 'サンプルカウンター罠','Trap15', 'テキスト5', 'text15' ),
(16, 6, 'サンプル永続魔法','Trap16', 'テキスト6', 'text16' ),
(17, 7, 'サンプル永続罠','Trap17', 'テキスト7', 'text17' ),
-- モンスターカード
(18, 8, 'もけもけ', 'Mokey Mokey', '天使のはみだし者', 'An outcast angel.'),
(19, 9, 'エクス・ライゼオル', 'Ext Ryzeal', 'テキスト１', 'text1'),
(20, 10, 'E・HERO アナザー・ネオス', 'Elemental HERO Neos Alius', 'テキスト１', 'text1'),
(21, 11, 'トゥーン・キャノン・ソルジャー', 'Toon Cannon Soldier', 'テキスト１', 'text1'),
(22, 12, 'イルミラージュ', 'Al-Lumiraj', 'テキスト１', 'text1'),
(23, 13, '運命の戦車', 'Fortune Chariot', 'テキスト１', 'text1'),
(24, 14, 'カオスポッド', 'Morphing Jar #2', 'テキスト１', 'text1'),
(25, 15, '砂塵の悪霊', 'Dark Dust Spirit', 'テキスト１', 'text1'),
(26, 16, 'シャドール・ファルコン', 'Shaddoll Falco', 'テキスト１', 'text1'),
(27, 17, 'カルボナーラ戦士', 'Karbonala Warrior', 'テキスト１', 'text1'),
(28, 18, '告死聖徒ルシエラーゴ', 'Azamina Mu Rcielago', 'テキスト１', 'text1'),
(29, 19, 'クラブ・タートル', 'Crab Turtle', 'テキスト１', 'text1'),
(30, 20, 'ウォーターリヴァイアサン＠イグニスター', 'Water Leviathan @Ignister', 'テキスト１', 'text1'),
(31, 21, 'ジェムナイト・パール', 'Gem-Knight Pearl', 'テキスト１', 'text1'),
(32, 22, 'ギガンティック・スプライト', 'Gigantic Spright', 'テキスト１', 'text1'),
(33, 23, 'スクラップ・デスデーモン', 'Scrap Archfiend', 'テキスト１', 'text1'),
(34, 24, 'Uk－P.U.N.K.アメイジング・ドラゴン', 'Ukiyoe-P.U.N.K. Amazing Dragon', 'テキスト１', 'text1'),
(35, 25, 'ガーデン・ローズ・フローラ', 'Garden Rose Flora', 'テキスト１', 'text1'),
(36, 26, 'ドラコニアの翼竜騎兵', 'Sky Dragoons of Draconia', 'テキスト１', 'Text2'),
(37, 27, 'アモルファージ・ルクス', 'Amorphage Lechery', 'テキスト１', 'text2'),
(38, 28, '覇王龍ズァーク', 'Supreme King Z-ARC', 'テキスト１', 'text2'),
(39, 29, 'DDD赦俿王デス・マキナ', 'D/D/D Deviser King Deus Machinex', 'テキスト１', 'text2'),
(40, 30, 'LANフォリンクス', 'LANphorhynchus', 'テキスト１', 'text1'),
(41, 31, 'アティプスの蟲惑魔', 'Traptrix Atypus', 'テキスト１', 'text1')
;

insert into
spells(card_id, spell_type_id)
values
(2, 1),
(4, 5),
(6, 2)
;

insert into
traps(card_id, trap_type_id)
values
(3, 1),
(5, 3),
(7, 2)
;

insert into
monsters(card_id, race_id, attribute_id, attack, defense, level, type_ids)
values
(1, 1, 2, 2000, 0, 4, ARRAY[1, 3]),
(8, 9, 1, 300, 100, 1, ARRAY[1]),
(9, 16, 1, 500, 2000, 4, ARRAY[2]),
(10, 4, 1, 1900, 1300, 4, ARRAY[6, 2]),
(11, 20, 2, 1400, 1300, 4, ARRAY[3, 2]),
(12, 23, 6, 1600, 1000, 3, ARRAY[7, 2]),
(13, 9, 6, 1000, 2000, 6, ARRAY[5, 2]),
(14, 18, 3, 800, 700, 3, ARRAY[8, 2]),
(15, 3, 3, 2200, 1800, 6, ARRAY[4, 2]),
(16, 1, 2, 600, 1400, 2, ARRAY[8, 7, 2]),
(17, 4, 3, 1500, 1200, 4, '{}'::integer[]),
(18, 25, 2, 2000, 2400, 6, ARRAY[2]),
(19, 15, 4, 2550, 2500, 8, '{}'::integer[]),
(20, 24, 4, 2300, 2000, 7, ARRAY[2]),
(21, 18, 3, 2600, 1900, 4, '{}'::integer[]),
(22, 17, 2, 1600, 1600, 2, ARRAY[2]),
(23, 8, 3, 2700, 1800, 7, '{}'::integer[]),
(24, 14, 6, 3000, 2800, 11, ARRAY[2]),
(25, 19, 1, 800, 1600, 5, ARRAY[7, 2]),
(26, 7, 6, 2200, 200, 5, ARRAY[1]),
(27, 2, 3, 1350, 0, 2, ARRAY[2]),
(28, 2, 2, 4000, 4000, 12, ARRAY[2]),
(29, 8, 2, 3000, 3000, 8, ARRAY[2]),
(30, 24, 1, 1200, 0, 0, '{}'::integer[]),
(31, 10, 3, 1800, 0, 0, ARRAY[2])
;

insert into fusion_monsters(card_id) values (17), (18), (28);
insert into ritual_monsters(card_id) values (19), (20);
insert into xyz_monsters(card_id) values (21), (22), (29);
insert into synchro_monsters(card_id) values (23), (24), (25);
insert into pendulum_monsters(card_id, scale, pendulum_text_ja, pendulum_text_en) values
(26, 7, 'テキスト２', 'Text2'),
(27, 5, 'テキスト２', 'text2'),
(28, 1, 'テキスト２', 'text2'),
(29, 10, 'テキスト２', 'text2');

insert into link_monsters(card_id, link_marker) values
(30, 40),
(31, 84);

COMMIT;