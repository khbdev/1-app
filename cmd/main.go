package main

import (
	"1-loyiha/handler"
	"1-loyiha/repostory"
	"1-loyiha/service"

	"github.com/gin-gonic/gin"
)


func main(){

	r := gin.Default()

		repo := repostory.NewUserRepostory()
	svc := service.NewUserService(repo)
	handler := handler.NewUserHandler(svc)

	handler.RegisterRoutes(r)

	r.Run(":8081")
	
}