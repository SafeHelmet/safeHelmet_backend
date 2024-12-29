package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

// list of all worksites routes
func RouteWorksites(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllWorksites)
	r.GET("/:worksite_id", controllers.GetWorksiteDetails)
	r.GET("/:worksite_id/workers", controllers.GetWorkersInWorksite)
	r.GET("/:worksite_id/readings", controllers.GetWorksiteReadings)

	r.POST("/assing-worker", controllers.AssignWorkerToWorksite)
}
