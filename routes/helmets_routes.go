package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteHelmets(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllHelmets)
	r.GET("/:helmet-id", controllers.GetHelmetDetails)
	r.GET("/mac-address/:mac-address", controllers.GetHelmetId)
	r.GET("/:helmet-id/readings", controllers.GetHelmetReadings)
	r.GET("/helmet-categories", controllers.GetHelmetCategories)
	r.GET("/helmet-categories/:category-id", controllers.GetHelmetCategoryDetails)

	r.GET("/:helmet-id/attendance", controllers.GetHelmetAttendance)

	r.PUT("/:helmet-id", controllers.UpdateHelmet)

	r.POST("", controllers.CreateHelmet)

	r.DELETE("/:helmet-id", controllers.DeleteHelmet)
}
