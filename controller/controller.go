package controller

import (
	"naive-server/api"
	"naive-server/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SignIn(c *gin.Context) {
	var req api.SignInReq
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "Could not bind json",
		})
		return
	}
	resp, err := model.SignIn(&req)
	if err != nil {
		logrus.Error(err)
		c.JSON(100, gin.H{
			"code": 100,
			"msg":  "invalid username or password",
		})
		return
	} else {
		c.JSON(200, resp)
		return
	}
}

func SignUp(c *gin.Context) {
	var req api.SignUpReq
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "Could not bind json",
		})
		return
	}
	resp, err := model.SignUp(&req)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal Server Error",
		})
		return
	} else {
		c.JSON(200, resp)
		return
	}
}

func CheckIn(c *gin.Context) {
	var req api.CheckInReq
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error(err)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "Could not bind json",
		})
		return
	}
	resp, err := model.CheckIn(&req)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Internal Server Error",
		})
		return
	} else {
		c.JSON(200, resp)
		return
	}
}
