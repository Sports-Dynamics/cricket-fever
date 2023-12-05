
-- +migrate Up
CREATE TABLE cricket_teams (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(50) NOT NULL,
    country VARCHAR(50) NOT NULL
);

CREATE INDEX idx_cricket_teams_team_name ON cricket_teams (team_name);
CREATE INDEX idx_cricket_teams_country ON cricket_teams (country);

-- +migrate Down
DROP Table cricket_teams;