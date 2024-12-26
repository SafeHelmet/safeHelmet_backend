package routes

import (
	"net/http"

	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func DeclareRoutes(r *gin.Engine) {

	// Definisci gli endpoint
	r.GET("/worksites/:worksiteId/workers", controllers.GetWorkers)
	r.GET("/workers/:workerId", controllers.GetWorkerDetails)
	r.GET("/workers/:workerId/worksites/:worksiteId/readings", controllers.GetReadings)
	r.GET("/worksites/:worksiteId/readings", controllers.GetWorksiteReadings)
	r.GET("/worksites/:worksiteId/readings/anomalous", controllers.GetAnomalousReadings)
	r.POST("/worksites/:worksiteId/workers", controllers.AssignWorkerToWorksite)

	// test
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test successful"})
	})

}
