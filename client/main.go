package main

import (
	"log"
	"parky/client/controller"
	"parky/client/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/api/authenticate_user", controller.AuthenticateUser)
	r.POST("/api/check_token", middlewares.TokenAuthMiddleware(), controller.CheckToken)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
