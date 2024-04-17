package repo

import (
	"github.com/awoyai/gin-temp/global"
	"github.com/awoyai/gin-temp/model"
)

type GameWhitelistRepo struct{}

func (repo *GameWhitelistRepo) CreateGameWhitelist(game *model.Game) error {
	return global.DB.Where("game_id = ?", game.GameID).FirstOrCreate(&game).Error
}

func (repo *GameWhitelistRepo) UpdateGameWhitelist(id uint, game *model.Game) error {
	return global.DB.Model(&model.Game{}).Where("id = ?", game.ID).Updates(&game).Error
}

func (repo *GameWhitelistRepo) DeleteGame(id uint) error {
	return global.DB.Delete(&model.Game{}, id).Error
}

func (repo *GameWhitelistRepo) ListGameWhitelist(filter *model.GameWhitelistFilter) ([]*model.Game, error) {
	var res []*model.Game
	db := global.DB
	if filter.GameHost != "" {
		db = db.Where("game_host = ?", filter.GameHost)
	}
	if filter.Email != "" {
		db = db.Where("email = ?", filter.Email)
	}
	if filter.GameID != "" {
		db = db.Where("game_id = ?", filter.GameID)
	}
	if filter.PageNum != 0 && filter.PageSize != 0 {
		db = db.Count(&filter.TotalCount).Offset(int((filter.PageNum - 1) * filter.PageSize)).Limit(int(filter.PageSize))
	}
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}
	if filter.PageNum != 0 && filter.PageSize != 0 {
		filter.TotalPage = (filter.TotalCount + filter.PageSize - 1) / filter.PageSize
	}
	return res, nil
}
