package routes

import (
	"net/http"

	"safecap_backend/controllers"

	"github.com/gin-gonic/gin"
)

func DeclareRoutes(r *gin.Engine) {

	// ALL
	r.GET("api/v1/worksites", controllers.GetAllWorksites)
	r.GET("api/v1/workers", controllers.GetAllWorkers)
	r.GET("api/v1/helmets", controllers.GetAllHelmets)
	r.GET("api/v1/helmetcategories", controllers.GetAllHelmetCategories)
	r.GET("api/v1/specializations", controllers.GetAllSpecializations)
	r.GET("api/v1/readings", controllers.GetAllReadings)
	r.GET("api/v1/workerworksiteassignments", controllers.GetAllWorkerWorksiteAssignments)
	r.GET("api/v1/worksitebossassignments", controllers.GetAllWorksiteBossAssignments)

	r.GET("api/v1/worksites/:worksiteId/workers", controllers.GetWorkers)
	r.GET("api/v1/workers/:workerId", controllers.GetWorkerDetails)
	r.GET("api/v1/workers/:workerId/worksites/:worksiteId/readings", controllers.GetReadings)
	r.GET("api/v1/worksites/:worksiteId/readings", controllers.GetWorksiteReadings)
	r.GET("api/v1/worksites/:worksiteId/readings/anomalous", controllers.GetAnomalousReadings)
	r.POST("api/v1/worksites/:worksiteId/workers", controllers.AssignWorkerToWorksite)

	// test
	r.GET("api/v1/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test successful"})
	})

}
