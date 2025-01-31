package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeclareRoutes(r *gin.Engine) {

	v1 := r.Group("api/v1")
	{
		worksites := v1.Group("/worksites")
		{
			RouteWorksites(worksites)
		}

		workers := v1.Group("/workers")
		{
			RouteWorkers(workers)
		}

		helmets := v1.Group("/helmets")
		{
			RouteHelmets(helmets)
		}

		readings := v1.Group("/readings")
		{
			RouteReadings(readings)
		}

		polling := v1.Group("/polling")
		{
			RoutePolling(polling)
		}
	}

	// test
	r.GET("api/v1/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test successful"})
	})

}
