package controller

import (
	"log"
	"net/http"
	"parky/client/token"
	"parky/client/utils"
	parkingpb "parky/proto"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func AuthenticateUser(c *gin.Context) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	var data map[string]interface{}
	err = c.BindJSON(&data)
	if err != nil {
		glog.Error(err)
	}
	if data["mobile"] == nil || !utils.ValidatePhone(data["mobile"].(string)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid request",
		})
		return
	}
	client := parkingpb.NewAuthenticationClient(conn)
	req := &parkingpb.AuthenticateUserRequest{Username: data["mobile"].(string)}
	res, err := client.AuthenticateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var tokenData token.AuthDetails
	tokenData.UserID = res.UserId
	token, loginErr := token.Authorize.SignIn(tokenData)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "error",
			"message": "Please try to login later",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func CheckToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"userId": c.MustGet("UserID").(string),
	})
}
