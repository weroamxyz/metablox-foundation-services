package controllers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/gin-gonic/gin"
)

//return all miners from DB. Most likely deprecated by the miner functions in metablox_staking
func GetMinerList(c *gin.Context) ([]models.MinerInfo, error) {
	minerList, err := dao.GetMinerList()
	if err != nil {
		return nil, err
	}

	return minerList, nil
}
