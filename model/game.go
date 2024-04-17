package model

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model/common"
)

type Game struct {
	global.MODEL
	Name         string `json:"name"`
	GameID       string `json:"game_id"`
	Domain       string `json:"domain"`
	DomainType   int    `json:"domain_type"`
	LoginType    string `json:"login_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_scret"`
	Scope        string `json:"scope"`
	FrontConfig  string `json:"front_config"`
	AuthUri      string `json:"auth_uri"`
	FaceUri      string `json:"face_uri"`
}

type GameFilter struct {
	Name   string
	GameID string
	Domain string
	common.PageInfo
}
