package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteReadings(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllReadings)
	r.GET("/:reading-id", controllers.GetReadingDetails)
	r.GET("/:reading-id/worker", controllers.GetReadingWorker)
	r.GET("/:reading-id/worksite", controllers.GetReadingWorksite)

	r.POST("", controllers.CreateReading)

	r.PUT("/:reading-id", controllers.UpdateReading)

	r.DELETE("/:reading-id", controllers.DeleteReading)
}
