package model

import "github.com/d-chosen-one/tgb/internal/enums"

type GameSettings struct {
	PlayerCount int               `json:"playerCount"`
	GameField   enums.EFieldSizes `json:"gameField"`
	PlayerName  string            `json:"playerName"`
	PlayerRace  enums.ERaces      `json:"playerRace"`
}

type Game struct {
	Id         string         `json:"id" bson:"_id"`
	PlayerName string         `json:"playerName" bson:"player_name"`
	PlayerRace enums.ERaces   `json:"playerRace" bson:"player_race"`
	Fields     []Field        `json:"fields" bson:"fields"`
	Races      []enums.ERaces `json:"races" bson:"races"`
}
