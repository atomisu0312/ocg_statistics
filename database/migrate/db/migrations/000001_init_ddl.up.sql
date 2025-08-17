CREATE TABLE spell_types (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE trap_types (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE monster_types (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE races (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE attributes (
    id INT PRIMARY KEY,
    name_ja VARCHAR(255),
    name_en VARCHAR(255),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE cards (
    id BIGSERIAL PRIMARY KEY,
    neuron_id BIGINT UNIQUE,
    ocg_api_id BIGINT UNIQUE,
    name_ja VARCHAR(255),
    name_en VARCHAR(255),
    card_text_ja TEXT,
    card_text_en TEXT,
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
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
    FOREIGN KEY (attribute_id) REFERENCES attributes(id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE ritual_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE xyz_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE synchro_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE fusion_monsters (
    card_id BIGINT PRIMARY KEY,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE link_monsters (
    card_id BIGINT PRIMARY KEY,
    link_marker INT,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE pendulum_monsters (
    card_id BIGINT PRIMARY KEY,
    scale INT,
    pendulum_text_ja TEXT,
    pendulum_text_en TEXT,
    FOREIGN KEY (card_id) REFERENCES monsters(card_id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE spells (
    card_id BIGINT PRIMARY KEY,
    spell_type_id INT,
    FOREIGN KEY (card_id) REFERENCES cards(id),
    FOREIGN KEY (spell_type_id) REFERENCES spell_types(id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);

CREATE TABLE traps (
    card_id BIGINT PRIMARY KEY,
    trap_type_id INT,
    FOREIGN KEY (card_id) REFERENCES cards(id),
    FOREIGN KEY (trap_type_id) REFERENCES trap_types(id),
    dataowner VARCHAR(255) DEFAULT 'system',
    regist_date TIMESTAMP DEFAULT current_timestamp,
    enable_start_date TIMESTAMP DEFAULT '1970-01-01 00:00:00',
    enable_end_date TIMESTAMP DEFAULT '9999-12-31 23:59:59',
    version BIGINT DEFAULT 1
);