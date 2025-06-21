package main

import (
"log"
"safecap_backend/API"
"safecap_backend/config"
"safecap_backend/controllers"
"safecap_backend/routes"
"safecap_backend/utils"
"safecap_backend/models"

)

func main() {
db := config.ConnectToDB(config.GetDSN())
config.Migrate(db)
// TODO scelta discutibile
controllers.InitDatabase(db)



var count int64
qCount :=db.Model(&models.HelmetCategory{}).Count(&count)
log.Println("QUERY COUNT RETURNS: ", qCount)
if err := db.Model(&models.HelmetCategory{}).Count(&count); err == 0 {
log.Println("Database need to be seeded!")
utils.DeleteTables(db)
err := utils.SeedDatabase(db)
if err != nil {
log.Fatal("Failed to seed database:", err)
}
log.Println("Database seeded successfully")
}

// Avvia il scheduler per la chiamata API in una goroutine
go API.StartAPICallScheduler(db)

// Crea un router Gin
r := config.SetupRouter()
routes.DeclareRoutes(r)

// Avvia il server sulla porta 8080
if err := r.Run(":8080"); err != nil {
log.Fatal(err)
}
}
