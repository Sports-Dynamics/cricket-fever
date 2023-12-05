
-- +migrate Up
CREATE TABLE team_players (
    player_id INTEGER NOT NULL REFERENCES cricket_players(player_id),
    team_id INTEGER NOT NULL REFERENCES cricket_teams(team_id),
    joining_date DATE NOT NULL,
    PRIMARY KEY (player_id, team_id, joining_date)
);

-- +migrate Down
DROP Table team_players;
