package request

type CreatePlayerRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	PlayerID string `validate:"required,min=1,max=200" json:"player_id"`
	LevelID  int    `validate:"required" json:"level_id"`
}

type UpdatePlayerRequest struct {
	ID       int    `validate:"required" json:"id"`
	Name     string `validate:"required,max=200,min=1" json:"name"`
	PlayerID string `validate:"required,min=1,max=200" json:"player_id"`
	LevelID  int    `validate:"required" json:"level_id"`
}
