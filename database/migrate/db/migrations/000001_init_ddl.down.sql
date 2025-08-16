-- created by following command
-- migrate create -ext sql -dir db/migrations -seq init_ddl

DROP TABLE IF EXISTS xyz_monsters;
DROP TABLE IF EXISTS synchro_monsters;
DROP TABLE IF EXISTS fusion_monsters;
DROP TABLE IF EXISTS link_monsters;
DROP TABLE IF EXISTS pendulum_monsters;
DROP TABLE IF EXISTS monsters;

DROP TABLE IF EXISTS spells;
DROP TABLE IF EXISTS traps;

DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS spell_types;
DROP TABLE IF EXISTS trap_types;
DROP TABLE IF EXISTS monster_types;
DROP TABLE IF EXISTS races;
DROP TABLE IF EXISTS attributes;