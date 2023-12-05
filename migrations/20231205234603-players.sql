
-- +migrate Up
CREATE TABLE cricket_players (
    player_id SERIAL PRIMARY KEY,
    player_name VARCHAR(50) NOT NULL,
    fielding_position VARCHAR(50) NOT NULL,
    batting_position INTEGER NOT NULL,
    bowler_type VARCHAR(50) NOT NULL
);

CREATE INDEX idx_cricket_players_player_name ON cricket_players (player_name);
CREATE INDEX idx_cricket_players_fielding_position ON cricket_players (fielding_position);
CREATE INDEX idx_cricket_players_batting_position ON cricket_players (batting_position);
CREATE INDEX idx_cricket_players_bowler_type ON cricket_players (bowler_type);

-- +migrate Down
DROP Table cricket_players;