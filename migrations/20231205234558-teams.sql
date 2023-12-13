
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE cricket_teams (
    team_id SERIAL PRIMARY KEY,
    team_uuid UUID DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    team_name VARCHAR(49) UNIQUE NOT NULL,
    team_country VARCHAR(49) NOT NULL,
    team_state VARCHAR(49) NOT NULL,
    team_city VARCHAR(49) NOT NULL,
    captain_id INT ,
    vice_captain_id INT ,
    team_coach_id INT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_cricket_teams_team_name ON cricket_teams (team_name);
CREATE INDEX idx_cricket_teams_country ON cricket_teams (team_country);


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

CREATE TABLE cricket_players (
    player_id SERIAL PRIMARY KEY,
    player_uuid UUID UNIQUE NOT NULL,
    player_name VARCHAR(50) UNIQUE NOT NULL,
    player_email VARCHAR(50) UNIQUE NOT NULL,
    player_mobile INT UNIQUE NOT NULL,
    player_picture BYTEA ,
    player_role cricket_role NOT NULL , 
    batting_positions cricket_batting_position NOT NULL,
    bowler_types cricket_bowler_type NOT NULL, 
    fielding_positions cricket_fielding_position NOT NULL,
    joining_date DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    team_id INTEGER NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (team_id) REFERENCES cricket_teams (team_id)
);

CREATE INDEX idx_cricket_players_player_uuid ON cricket_players (player_uuid);
CREATE INDEX idx_cricket_players_player_name ON cricket_players (player_name);
CREATE INDEX idx_cricket_players_batting_position ON cricket_players (batting_positions);
CREATE INDEX idx_cricket_players_bowler_type ON cricket_players (bowler_types);
CREATE INDEX idx_cricket_players_fielding_position ON cricket_players (fielding_positions);

CREATE TABLE team_players(
    player_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL,
    joining_date DATE NOT NULL,
    FOREIGN KEY (player_id) REFERENCES cricket_players (player_id),
    FOREIGN KEY (team_id) REFERENCES cricket_teams (team_id),
     PRIMARY KEY (player_id, team_id)
);


-- +migrate Down
DROP TABLE team_players;
DROP Table cricket_teams;
DROP Table cricket_players;

DROP TYPE cricket_role;
DROP TYPE cricket_batting_position;
DROP TYPE cricket_bowler_type;
DROP TYPE cricket_fielding_position;
