package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/test/controllers"
	_ "github.com/test/docs"
	"github.com/test/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDb()
	initializers.InsertFixedSubscription(initializers.DB)
}
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//CRUDL
	r.POST("/subs", controllers.AddSubs)
	r.GET("/subs", controllers.GetListOfIDs) // query controllers.ListOfIDs with filters (user_id and service name)
	r.GET("/subs/:id", controllers.GetByID)
	r.PUT("/subs/:id", controllers.UpdateByID)
	r.DELETE("/subs/:id", controllers.DeleteByID)

	//FILTERED
	r.GET("/subs/summary", controllers.GetSubSummary)
	r.Run()
}
