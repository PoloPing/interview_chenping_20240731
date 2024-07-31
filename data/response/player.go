package response

type PlayerResponse struct {
	ID            uint          `json:"id"`
	Name          string        `json:"name"`
	PlayerID      string        `json:"player_id"`
	LevelResponse LevelResponse `json:"level"`
}
