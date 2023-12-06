

CREATE TYPE cricket_role  AS ENUM(
'Batsman',
'Bowler',
'Batsman AllRounder',
'Bowling AllRounder',
'Wicket Keeper'
);

CREATE TYPE  cricket_batting_position AS ENUM (
'Opening Batsman',
'Top Order',
'Middle Order',
'Lower Middle Order',
'Tailender'
);

CREATE TYPE cricket_bowler_type AS ENUM (
'Fast Bowler',
'Medium Fast Bowler',
'Medium Pacer',
'Spin Bowler',
'Leg Spinner',
'Off Spinner',
'ChinaMen'
);


CREATE TYPE  cricket_fielding_position AS ENUM (
'Slip',
'Gully',
'Third Man',
'Fine Leg',
'Deep Square Leg',
'Deep Midwicket',
'Long On',
'Long Off',
'Cover',
'Extra Cover',
'Point',
'Square Leg',
'Mid On',
'Mid Off',
'Midwicket',
'Short Leg',
'Leg Slip',
'Short Fine Leg'
);

-- +migrate Up

CREATE TABLE cricket_players (
    player_id SERIAL PRIMARY KEY,
    player_name VARCHAR(50) UNIQUE NOT NULL,
    player_email VARCHAR(50) UNIQUE NOT NULL,
    player_mobile INT UNIQUE NOT NULL,
    player_picture BYTEA ,
    player_role cricket_role , 
    batting_positions cricket_batting_position,
    bowler_types cricket_bowler_type, 
    fielding_positions cricket_fielding_position
);

CREATE INDEX idx_cricket_players_player_name ON cricket_players (player_name);
CREATE INDEX idx_cricket_players_batting_position ON cricket_players (batting_positions);
CREATE INDEX idx_cricket_players_bowler_type ON cricket_players (bowler_types);
CREATE INDEX idx_cricket_players_fielding_position ON cricket_players (fielding_positions);

-- +migrate Down
DROP Table cricket_players;