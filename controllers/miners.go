package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/dao"
	"github.com/metabloxDID/did"
	"github.com/metabloxDID/errval"
	"github.com/metabloxDID/models"
)

func GetMinerList(c *gin.Context) ([]models.MinerInfo, error) {
	didString := "did:metablox:" + c.Param("did")
	valid := did.IsDIDValid(did.SplitDIDString(didString))
	if !valid {
		return nil, errval.ErrDIDFormat
	}

	minerList, err := dao.GetMinerList()
	if err != nil {
		return nil, err
	}

	return minerList, nil
}
