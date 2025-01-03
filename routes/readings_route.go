package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteReadings(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllReadings)
	r.GET("/:reading-id", controllers.GetReadingDetails)
}
