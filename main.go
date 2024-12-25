package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Simulazione dei dati
var workers = map[string][]string{
	"worksite1": {"worker1", "worker2"},
	"worksite2": {"worker3", "worker4"},
}

var workerDetails = map[string]map[string]interface{}{
	"worker1": {"specialization": "electrician", "assignments": []string{"worksite1"}},
	"worker2": {"specialization": "plumber", "assignments": []string{"worksite1"}},
}

var readings = map[string]map[string]map[string]interface{}{
	"worker1": {
		"worksite1": {"2024-12-25": []string{"reading1", "reading2"}},
	},
	"worker2": {
		"worksite1": {"2024-12-25": []string{"reading3", "reading4"}},
	},
}

// Main handler function to initialize the server
func main() {
	// Crea un router Gin
	r := gin.Default()

	// Definisci gli endpoint
	r.GET("/worksites/:worksiteId/workers", getWorkers)
	r.GET("/workers/:workerId", getWorkerDetails)
	r.GET("/workers/:workerId/worksites/:worksiteId/readings", getReadings)
	r.GET("/worksites/:worksiteId/readings", getWorksiteReadings)
	r.GET("/worksites/:worksiteId/readings/anomalous", getAnomalousReadings)
	r.POST("/worksites/:worksiteId/workers", assignWorkerToWorksite)

	// test
	r.GET("/test", test)

	// Avvia il server sulla porta 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "paolo bellavista gran frocio",
	})
}

// Elenco dei Lavoratori in un Cantiere
func getWorkers(c *gin.Context) {
	worksiteId := c.Param("worksiteId")

	// Verifica se il worksiteId esiste
	if workers, exists := workers[worksiteId]; exists {
		c.JSON(http.StatusOK, gin.H{
			"worksiteId": worksiteId,
			"workers":    workers,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Worksite not found"})
	}
}

// Dettagli di un Lavoratore
func getWorkerDetails(c *gin.Context) {
	workerId := c.Param("workerId")

	// Verifica se il lavoratore esiste
	if details, exists := workerDetails[workerId]; exists {
		c.JSON(http.StatusOK, gin.H{
			"workerId":       workerId,
			"specialization": details["specialization"],
			"assignments":    details["assignments"],
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
	}
}

// Letture di un Casco per un Lavoratore in un Cantiere
func getReadings(c *gin.Context) {
	workerId := c.Param("workerId")
	worksiteId := c.Param("worksiteId")
	date := c.DefaultQuery("date", "2024-12-25") // Data con valore di default

	// Verifica se esistono letture per quel lavoratore, cantiere e data
	if workerReadings, exists := readings[workerId]; exists {
		if worksiteReadings, exists := workerReadings[worksiteId]; exists {
			if readings, exists := worksiteReadings[date]; exists {
				c.JSON(http.StatusOK, gin.H{
					"workerId":   workerId,
					"worksiteId": worksiteId,
					"date":       date,
					"readings":   readings,
				})
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Readings not found"})
}

// Letture di un Cantiere in un Intervallo Temporale
func getWorksiteReadings(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	from := c.DefaultQuery("from", "2024-12-01 00:00:00")
	to := c.DefaultQuery("to", "2024-12-01 23:59:59")

	// Qui simuliamo che ci siano letture per il worksite, in una vera implementazione ci sarebbe una query al DB
	c.JSON(http.StatusOK, gin.H{
		"worksiteId": worksiteId,
		"from":       from,
		"to":         to,
		"readings":   []string{"reading1", "reading2"}, // Letture fittizie
	})
}

// Letture Anomale di un Cantiere
func getAnomalousReadings(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	// Simuliamo letture anomale per il worksite
	c.JSON(http.StatusOK, gin.H{
		"worksiteId":        worksiteId,
		"anomalousReadings": []string{"anomalousReading1", "anomalousReading2"},
	})
}

// Assegna un Lavoratore a un Cantiere con un Casco Specifico
func assignWorkerToWorksite(c *gin.Context) {
	worksiteId := c.Param("worksiteId")
	var requestBody struct {
		WorkerId int `json:"workerId"`
		HelmetId int `json:"helmetId"`
	}

	// Legge il corpo della richiesta (in JSON)
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simula l'assegnazione del lavoratore al cantiere
	c.JSON(http.StatusOK, gin.H{
		"message":    "Worker assigned successfully",
		"workerId":   requestBody.WorkerId,
		"helmetId":   requestBody.HelmetId,
		"worksiteId": worksiteId,
	})
}
