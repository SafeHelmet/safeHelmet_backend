package main

import (
	"log"
	"os"
	"safecap_backend/config"
	"safecap_backend/routes"
	"safecap_backend/utils"
)

func main() {
	var err error
	toseed := "false"

	// Verifica se è stato passato un argomento da riga di comando
	if len(os.Args) >= 2 {
		toseed = os.Args[1]
	}

	db := config.ConnectToDB(config.GetDSN())
	config.Migrate(db)

	if toseed == "true" {
		utils.DeleteTables(db)
		err = utils.SeedDatabase(db)
		if err != nil {
			log.Fatal("Failed to seed database:", err)
		}
		log.Println("Database seeded successfully")
	}

	// Crea un router Gin
	r := config.SetupRouter()
	routes.DeclareRoutes(r)

	// Avvia il server sulla porta 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
