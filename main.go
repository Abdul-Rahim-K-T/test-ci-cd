package main

import (
	"github.com/Abdul-Rahim-K-T/pipeline/database"
	"github.com/Abdul-Rahim-K-T/pipeline/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	database.InitDB()
	router:=gin.Default()
	routes.Router(router)
	router.Run("localhost:8080")
}