
-- +migrate Up
CREATE TABLE cricket_teams (
    team_id SERIAL PRIMARY KEY,
    team_uuid UUID UNIQUE NOT NULL,
    team_name VARCHAR(50) UNIQUE NOT NULL,
    team_country VARCHAR(50) NOT NULL,
    team_state VARCHAR(50) NOT NULL,
    team_city VARCHAR(50) NOT NULL,
    captain_id INT ,
    vice_captain_id INT ,
    team_coach_id INT 
);

CREATE INDEX idx_cricket_teams_team_name ON cricket_teams (team_name);
CREATE INDEX idx_cricket_teams_country ON cricket_teams (team_country);

-- +migrate Down
DROP Table cricket_teams;