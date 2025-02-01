package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteHelmets(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllHelmets)
	r.GET("/:helmet-id", controllers.GetHelmetDetails)
	r.GET("/:helmet-uuid", controllers.GetHelmetId)
}
