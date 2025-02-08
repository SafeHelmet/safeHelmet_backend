package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteWorkers(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllWorkers)
	r.GET("/:worker-id", controllers.GetWorkerDetails)
	r.GET("/:worker-id/worksite", controllers.GetWorksiteOfWorker)

	r.GET("/:worker-id/attendance", controllers.GetWorkerAttendance)

	r.POST("", controllers.CreateWorker)
	r.POST("/attendance", controllers.CreateWorkerAttendance)

	r.PUT("/:worker-id", controllers.UpdateWorker)
	r.PUT("/attendance", controllers.UpdateWorkerAttendance)

	r.DELETE("/:worker-id", controllers.DeleteWorker)
}
