package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

// list of all worksites routes
func RouteWorksites(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllWorksites)
	r.GET("/:id", controllers.GetWorksiteDetails)
	r.GET("/:id/workers", controllers.GetWorkersInWorksite)
	r.GET("/:id/readings", controllers.GetWorksiteReadings)

	
}
