package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RoutePolling(r *gin.RouterGroup) {
	r.GET("/:helmet-id", controllers.CheckRecentAnomaly)
}
