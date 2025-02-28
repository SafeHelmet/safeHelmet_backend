package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteWorkers(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllWorkers)
	r.GET("/:worker-id", controllers.GetWorkerDetails)
	r.GET("/:worker-id/worksite", controllers.GetWorksiteOfWorker)
	r.GET("/:worker-id/readings", controllers.GetWorkerReadings)

	r.GET("/:worker-id/attendance", controllers.GetWorkerAttendance)
	r.GET("/:worker-id/attendance/last", controllers.GetLastWorkerAttendance)

	r.POST("/attendance", controllers.CreateWorkerAttendance)
	r.POST("", controllers.CreateWorker)

	r.PUT("/:worker-id", controllers.UpdateWorker)
	r.PUT("/attendance/:attendance-id", controllers.UpdateWorkerAttendance)

	r.DELETE("/:worker-id", controllers.DeleteWorker)
}
