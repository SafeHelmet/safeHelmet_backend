package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteWorkers(r *gin.RouterGroup) {
	r.GET("/bosses", controllers.GetAllBosses)
	r.GET("", controllers.GetAllWorkers)
	r.GET("/:id", controllers.GetWorkerDetails)

}
