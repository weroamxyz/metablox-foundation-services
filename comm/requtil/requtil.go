package requtil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ShouldBind[T any](c *gin.Context) (*T, error) {
	params := new(T)
	if err := c.ShouldBind(params); err != nil {
		return nil, err
	}
	return params, nil
}

func ShouldBindJSON[T any](c *gin.Context) (*T, error) {
	params := new(T)
	if err := c.ShouldBindJSON(params); err != nil {
		return nil, err
	}
	return params, nil
}

type UserInfo struct {
	UserType    string
	UserSub     string
	UserCompany string
}

func ShouldBindQuery[T any](c *gin.Context) (*T, error) {
	params := new(T)
	return params, c.ShouldBindQuery(params)
}

func ShouldQueryString(c *gin.Context, key string) (query string, err error) {
	if query = c.Query(key); query == "" {
		return query, fmt.Errorf("%s cannot be empty", key)
	}
	return query, nil
}

func ShouldQueryInt(c *gin.Context, key string) (queryInt int, err error) {
	if query := c.Query(key); query == "" {
		return 0, fmt.Errorf("%s cannot be empty", key)
	} else {
		if queryInt, err = strconv.Atoi(query); err != nil {
			return 0, err
		}
		return queryInt, nil
	}

}

func BindQueryMap(c *gin.Context) map[string]string {
	params := make(map[string]string)
	c.ShouldBindQuery(params)
	return params
}

func ShouldParam(c *gin.Context, key string) (param string, err error) {
	if param = c.Param(key); param == "" {
		return param, fmt.Errorf("%s cannot be empty", key)
	}
	return param, nil
}
