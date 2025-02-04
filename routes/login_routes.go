package routes

import (
	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func RouteLogin(r *gin.RouterGroup) {
	r.POST("", controllers.Login)
}
