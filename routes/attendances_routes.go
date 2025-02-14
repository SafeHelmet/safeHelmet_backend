package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteAttendances(r *gin.RouterGroup) {

	r.GET("", controllers.GetAllAttendances)
	r.GET("/:attendance_id", controllers.GetAttendanceDetails)

	r.GET("/attendance-details/:worker_id/:worksite_id/:helmet_id", controllers.GetLastAttendanceDetails)
}
