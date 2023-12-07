package teams

type CreateTeamRequestParams struct {
	Name            string `json:"team_name"`
	Country         string `json:"team_country"`
	State           string `json:"team_state"`
	City            string `json:"team_city"`
	CaptainUUID     string `json:"team_captain"`
	ViceCaptainUUID string `json:"team_vice_captain"`
	TeamCoachUUID   string `json:"team_coach"`
}

func (t *CreateTeamRequestParams) Validate() error {
	return nil
}
