package routes

import (
	"GoZapLearn/01_HelloZap/config"
	"GoZapLearn/01_HelloZap/controller"
	"GoZapLearn/01_HelloZap/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Middleware())
	router.GET("/Person", controller.GetPerson)
	router.POST("/Person", controller.CreatePerson)
	router.DELETE("/Person/:id", controller.DeletePerson)
	router.PUT("/Person/:id", controller.UpdatePerson)
	return router
}

func RunRouter(c *gin.Engine) {
	err := c.Run(config.ListenAt)
	if err != nil {
		panic(err)
	}
}
