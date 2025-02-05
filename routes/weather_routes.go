package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteWeather(r *gin.RouterGroup) {
	r.GET("/:worksite-id", controllers.GetAllWorksiteWeather)
	r.GET("/last/:worksite-id", controllers.GetLastWorksiteWeather)
	r.GET("", controllers.GetAllWeatherReadings)
}
