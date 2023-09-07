package model

import (
	"time"

	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
)

type GameWhitelist struct {
	global.MODEL
	GameId    string    `json:"game_id"`
	Email     string    `json:"email"`
	StartAt   time.Time `json:"start_at"`
	ExpiredAt time.Time `json:"expired_at"`
	GameHost  string    `json:"game_host"`
}

type GameWhitelistFilter struct {
	GameID   string
	GameHost string
	Email    string
	common.PageInfo
}

type WebWhitelist struct {
	global.MODEL
	Email     string    `json:"email"`
	StartAt   time.Time `json:"start_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type WebWhitelistFilter struct {
	Email string
	common.PageInfo
}
