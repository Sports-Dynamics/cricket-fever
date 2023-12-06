package teams 

type CreateTeamRequestParams struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (t *CreateTeamRequestParams) Validate() error {
	return nil
}
