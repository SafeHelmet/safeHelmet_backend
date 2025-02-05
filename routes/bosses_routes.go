package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteBosses(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllBosses)
	r.GET("/:boss-id", controllers.GetBossDetails)
	r.POST("", controllers.CreateBoss)

	r.PUT("/:boss-id", controllers.UpdateWorker)

	r.DELETE("/:boss-id", controllers.DeleteBoss)
}
