CREATE TABLE spell_types (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255)
);

CREATE TABLE trap_types (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255)
);

CREATE TABLE monster_types (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255)
);

CREATE TABLE races (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255)
);

CREATE TABLE attributes (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255)
);

CREATE TABLE cards (
    id BIGSERIAL PRIMARY KEY,
    neuron_id BIGINT UNIQUE,
    ocg_api_id BIGINT UNIQUE,
    name_ja VARCHAR(255),
    name_en VARCHAR(255),
    card_text_ja TEXT,
    card_text_en TEXT
);

CREATE TABLE monsters (
    card_id BIGINT PRIMARY KEY,
    race_id INT,
    attribute_id INT,
    attack INT,
    defense INT,
    level INT,
    type_ids INT[],
    FOREIGN KEY (card_id) REFERENCES cards(id),
    FOREIGN KEY (race_id) REFERENCES races(id),
    FOREIGN KEY (attribute_id) REFERENCES attributes(id)
);

CREATE TABLE ritual_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id)
);

CREATE TABLE xyz_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id)
);

CREATE TABLE synchro_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id)
);

CREATE TABLE fusion_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id)
);

CREATE TABLE link_monsters (
    card_id BIGINT PRIMARY KEY,
    link_marker INT,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id)
);

CREATE TABLE pendulum_monsters (
    card_id BIGINT PRIMARY KEY,
    scale INT,
    text_ja TEXT,
    text_en TEXT,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id)
);

CREATE TABLE spells (
    card_id BIGINT PRIMARY KEY,
    id INT,
    FOREIGN KEY (card_id) REFERENCES cards(id),
    FOREIGN KEY (id) REFERENCES spell_types(id)
);

CREATE TABLE traps (
    card_id BIGINT PRIMARY KEY,
    id INT,
    FOREIGN KEY (card_id) REFERENCES cards(id),
    FOREIGN KEY (id) REFERENCES trap_types(id)
);
