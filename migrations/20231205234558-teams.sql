-- Active: 1700067648591@@localhost@5432@postgres

-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE address_type AS (
    team_country VARCHAR(49) ,
    team_state VARCHAR(49) ,
    team_city VARCHAR(49)
);


CREATE TABLE cricket_teams (
    team_id SERIAL PRIMARY KEY,
    team_uuid UUID DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
    team_name VARCHAR(49) UNIQUE NOT NULL,
    team_address address_type ,
    captain_id INT ,
    vice_captain_id INT ,
    team_coach_id INT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX idx_cricket_teams_team_name ON cricket_teams (team_name);


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

CREATE TABLE cricket_grounds (
    ground_id SERIAL PRIMARY KEY,
    ground_name VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    ground_address address_type,
    capacity INT,
    established_year INT,
    listed_year INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create table for bookings
CREATE TABLE ground_bookings (
    booking_id SERIAL PRIMARY KEY,
    team_id INT REFERENCES cricket_teams(team_id),
    player_id INT REFERENCES cricket_players(player_id),
    ground_id INT REFERENCES cricket_grounds(ground_id),
    booking_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    purpose VARCHAR(255),
    status VARCHAR(50) DEFAULT 'Pending', -- Added booking status column
    CONSTRAINT valid_time_range CHECK (start_time < end_time)
);

-- +migrate StatementBegin
-- Create a function to be used in the trigger
CREATE OR REPLACE FUNCTION check_team_count()
RETURNS TRIGGER AS $$
BEGIN
    IF (
        SELECT COUNT(DISTINCT team_id) 
        FROM ground_bookings 
        WHERE ground_id = NEW.ground_id
    ) > 2 THEN
        RAISE EXCEPTION 'More than two teams cannot have the same ground_id';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- Create the trigger
CREATE TRIGGER enforce_team_count
BEFORE INSERT OR UPDATE ON ground_bookings
FOR EACH ROW EXECUTE FUNCTION check_team_count();


-- Create table for payments
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    booking_id INT REFERENCES ground_bookings(booking_id),
    amount DECIMAL(10, 2) NOT NULL,
    payment_date DATE NOT NULL,
    payment_status VARCHAR(50) DEFAULT 'Pending'
);


-- +migrate Down
DROP TABLE payments ; 
DROP TABLE ground_bookings;
DROP TABLE team_players;
DROP TABLE cricket_players;
DROP Table cricket_teams;
DROP TYPE cricket_role;
DROP TYPE cricket_batting_position;
DROP TABLE cricket_grounds;
DROP TYPE cricket_bowler_type;
DROP TYPE cricket_fielding_position;
DROP TYPE address_type;
