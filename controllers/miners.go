package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func GetMinerList(c *gin.Context) ([]models.MinerInfo, error) {
	minerList, err := dao.GetMinerList()
	if err != nil {
		return nil, err
	}

	return minerList, nil
}
