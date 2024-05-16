package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	//"net/http"
	"tour-service/handler"
	"tour-service/model"
	"tour-service/proto/tour"
	"tour-service/repository"
	"tour-service/service"

	//"github.com/gorilla/handlers"
	//"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	//connectionStr := "root:super@tcp(localhost:3306)/soa?charset=utf8mb4&parseTime=True&loc=Local"
	connectionStr := "root:root@tcp(tour_db:3306)/soa?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Equipment{}, &model.Checkpoint{}, &model.Tour{}, &model.TourReview{}, &model.TourProblem{})
	//database.Exec("INSERT IGNORE INTO tours VALUES ('aec7e123-243d-4a09-a289-75308ea5b7e6', 1, 'Naslov ture', 'Opis ture', '2024-03-19 12:00:00', 1, 'putanja/do/slike.jpg', 3, 2999, 'tag1', 10.5, 26.25, 42.08, 40.5)")
	return database
}
func startServer(tourHandler *handler.TourHandler, checkpointHandler *handler.CheckpointHandler, equipmentHandler *handler.EquipmentHandler, tourReviewHandler *handler.TourReviewHandler, tourProblemHandler *handler.TourProblemHandler) {
	//router := mux.NewRouter().StrictSlash(true)

	/*router.HandleFunc("/tour/{id}", tourHandler.Get).Methods("GET")
	router.HandleFunc("/tour", tourHandler.GetAll).Methods("GET")
	router.HandleFunc("/tour", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tour/{id}", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/tour/{id}", tourHandler.Delete).Methods("DELETE")
	router.HandleFunc("/toursByAuthorId/{authorId}", tourHandler.GetToursByAuthor).Methods("GET")

	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Get).Methods("GET")
	router.HandleFunc("/checkpoint", checkpointHandler.Create).Methods("POST")
	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Update).Methods("PUT")
	router.HandleFunc("/checkpoint/{id}", checkpointHandler.Delete).Methods("DELETE")
	router.HandleFunc("/checkpointByTourID/{id}", checkpointHandler.GetCheckpointsByTourID).Methods("GET")

	router.HandleFunc("/equipment/{id}", equipmentHandler.Get).Methods("GET")
	router.HandleFunc("/equipment", equipmentHandler.Create).Methods("POST")
	router.HandleFunc("/equipment/{id}", equipmentHandler.Update).Methods("PUT")
	router.HandleFunc("/equipment/{id}", equipmentHandler.Delete).Methods("DELETE")

	router.HandleFunc("/tourReview", tourReviewHandler.GetAll).Methods("GET")
	router.HandleFunc("/tourReview/{id}", tourReviewHandler.GetTourReviewsByTourID).Methods("GET")
	router.HandleFunc("/tourReview", tourReviewHandler.Create).Methods("POST")
	router.HandleFunc("/tourReview/{id}", tourReviewHandler.Update).Methods("PUT")
	router.HandleFunc("/tourReview/{id}", tourReviewHandler.Delete).Methods("DELETE")
	router.HandleFunc("/tourReview/average-grade/{id}", tourReviewHandler.GetAverageGradeForTour).Methods("GET")

	router.HandleFunc("/tourProblems/{id}", tourProblemHandler.GetTourProblemsForTourist).Methods("GET")
	router.HandleFunc("/tourProblem/{id}", tourProblemHandler.Get).Methods("GET")
	router.HandleFunc("/tourProblem", tourProblemHandler.Create).Methods("POST")
	router.HandleFunc("/tourProblem", tourProblemHandler.Update).Methods("PUT")
	router.HandleFunc("/tourProblem/{id}", tourProblemHandler.Delete).Methods("DELETE")*/

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	tour.RegisterTourServiceServer(grpcServer, tourHandler)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()

	/*allowedOrigins := handlers.AllowedOrigins([]string{"*"}) // Allow all origins
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
	log.Fatal(http.ListenAndServe(":8081", corsRouter))*/

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
	tourReviewRepo := &repository.TourReviewRepository{DatabaseConnection: database}
	tourProblemRepo := &repository.TourProblemRepository{DatabaseConnection: database}

	checkpointService := &service.CheckpointService{CheckpointRepo: checkpointRepo}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourReviewService := &service.TourReviewService{TourReviewRepo: tourReviewRepo}
	tourProblemService := &service.TourProblemService{TourProblemRepo: tourProblemRepo}

	checkpointHandler := &handler.CheckpointHandler{CheckpointService: checkpointService}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}
	tourHandler := &handler.TourHandler{TourService: tourService}
	tourReviewHandler := &handler.TourReviewHandler{TourReviewService: tourReviewService}
	tourProblemHandler := &handler.TourProblemHandler{TourProblemService: tourProblemService}

	startServer(tourHandler, checkpointHandler, equipmentHandler, tourReviewHandler, tourProblemHandler)
}
