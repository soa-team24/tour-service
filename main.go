package main

import (
	"log"
	"net/http"
	"tour-service/handler"
	"tour-service/model"
	"tour-service/repository"
	"tour-service/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "root:super@tcp(localhost:3306)/soa?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Equipment{}, &model.Checkpoint{}, &model.Tour{})

	return database
}

func startServer(tourHandler *handler.TourHandler, checkpointHandler *handler.CheckpointHandler, equipmentHandler *handler.EquipmentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tour/{id}", tourHandler.Get).Methods("GET")
	router.HandleFunc("/tour", tourHandler.GetAll).Methods("GET")
	router.HandleFunc("/tour", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tour/{id}", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/tour/{id}", tourHandler.Delete).Methods("DELETE")

	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Get).Methods("GET")
	router.HandleFunc("/checkpoint", checkpointHandler.Create).Methods("POST")
	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Update).Methods("PUT")
	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Delete).Methods("DELETE")

	router.HandleFunc("/equipment/{id}", equipmentHandler.Get).Methods("GET")
	router.HandleFunc("/equipment", equipmentHandler.Create).Methods("POST")
	router.HandleFunc("/equipment/{id}", equipmentHandler.Update).Methods("PUT")
	router.HandleFunc("/equipment/{id}", equipmentHandler.Delete).Methods("DELETE")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) // Allow all origins
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{
		"Content-Type",
		"Authorization",
		"X-Custom-Header",
	})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Apply CORS middleware to all routes
	corsRouter := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", corsRouter))

}

func main() {

	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	checkpointRepo := &repository.CheckpointRepository{DatabaseConnection: database}
	equipmentRepo := &repository.EquipmentRepository{DatabaseConnection: database}
	tourRepo := &repository.TourRepository{DatabaseConnection: database}

	checkpointService := &service.CheckpointService{CheckpointRepo: checkpointRepo}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	tourService := &service.TourService{TourRepo: tourRepo}

	checkpointHandler := &handler.CheckpointHandler{CheckpointService: checkpointService}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}
	tourHandler := &handler.TourHandler{TourService: tourService}

	startServer(tourHandler, checkpointHandler, equipmentHandler)

}
