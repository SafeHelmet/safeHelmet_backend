package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

// list of all worksites routes
func RouteWorksites(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllWorksites)
	r.GET("/:worksite-id", controllers.GetWorksiteDetails)
	r.GET("/:worksite-id/workers", controllers.GetWorkersInWorksite)
	r.GET("/:worksite-id/readings", controllers.GetWorksiteReadings)

	r.GET("/:worksite-id/attendance", controllers.GetWorksiteAttendance)

	r.POST("", controllers.CreateWorksite)
	r.POST("/assing-worker", controllers.AssignWorkerToWorksite)

	r.PUT("/:worksite-id", controllers.UpdateWorksite)

	r.DELETE("/:worksite-id", controllers.DeleteWorksite)
}
