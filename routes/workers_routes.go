package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteWorkers(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllWorkers)
	r.GET("/:worker-id", controllers.GetWorkerDetails)
	r.GET("/:worker-id/worksite", controllers.GetWorksiteOfWorker)

	r.POST("", controllers.CreateWorker)

	r.PUT("/:worker-id", controllers.UpdateWorker)

	r.DELETE("/:worker-id", controllers.DeleteWorker)
}
