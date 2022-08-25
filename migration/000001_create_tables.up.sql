CREATE TABLE IF NOT EXISTS games(
    id VARCHAR(36),
    active_player_id VARCHAR(36) NOT NULL,
    status VARCHAR(20) NOT NULL,
    turn INT NOT NULL,
    version INT,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS players(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    color VARCHAR(36) NOT NULL,
    turn_order INT NOT NULL,
    version INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS constructions(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    player_id VARCHAR(36) NOT NULL,
    type VARCHAR(20) NOT NULL,
    version INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS roads(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    player_id VARCHAR(36) NOT NULL,
    version INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dices(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    number INT NOT NULL,
    is_rolled BOOLEAN NOT NULL,
    version INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS achievements(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    player_id VARCHAR(36),
    type VARCHAR(20) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS resource_cards(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    player_id VARCHAR(36),
    type VARCHAR(20) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS development_cards(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    player_id VARCHAR(36),
    type VARCHAR(20) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS terrains(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    q INT NOT NULL,
    r INT NOT NULL,
    number INT NOT NULL,
    type VARCHAR(20) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS harbors(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    terrain_id VARCHAR(36) NOT NULL,
    Q INT,
    R INT,
    type VARCHAR(20) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (terrain_id) REFERENCES terrains(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS robbers(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    terrain_id VARCHAR(36) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (terrain_id) REFERENCES terrains(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS lands(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    construction_id VARCHAR(36),
    q INT,
    r INT,
    location VARCHAR(20) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (construction_id) REFERENCES constructions(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS paths(
    id VARCHAR(36),
    game_id VARCHAR(36) NOT NULL,
    road_id VARCHAR(36),
    q INT,
    r INT,
    location VARCHAR(20) NOT NULL,
    version INT,
    PRIMARY KEY (id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (road_id) REFERENCES roads(id) ON DELETE CASCADE
);


