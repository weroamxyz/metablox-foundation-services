package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/dao"
	"github.com/metabloxDID/models"
)

func GetMinerList(c *gin.Context) ([]models.MinerInfo, error) {
	minerList, err := dao.GetMinerList()
	if err != nil {
		return nil, err
	}

	return minerList, nil
}
