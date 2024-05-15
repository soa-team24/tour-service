package handler

import (
	"context"
	"fmt"
	"log"
	"soa/grpc/proto/tour"
	"strconv"
	"tour-service/dto"
	"tour-service/mapper"
	"tour-service/service"
)

type KeyProduct struct{}

type TourHandler struct {
	tour.UnimplementedTourServiceServer
	TourService        *service.TourService
	CheckpointService  *service.CheckpointService
	EquipmentService   *service.EquipmentService
	TourProblemService *service.TourProblemService
	TourReviewService  *service.TourReviewService
}

func (handler *TourHandler) GetAllTours(ctx context.Context, request *tour.GetAllRequest) (*tour.GetToursResponse, error) {
	tours, err := handler.TourService.GetAll()
	if err != nil {
		println("Database exception: ")
	}

	if tours == nil {
		return nil, fmt.Errorf("failed to get blogs")
	}
	protoTours := mapper.MapSliceToProtoTours(tours)
	response := &tour.GetToursResponse{
		Tours: protoTours,
	}
	return response, nil

}

func (handler *TourHandler) GetTour(ctx context.Context, request *tour.GetRequest) (*tour.TourResponse, error) {
	id := request.Id
	log.Printf("Tour sa id-em %s", id)
	modelTour, err := handler.TourService.Get(id)
	if err != nil {
		println("Database exception: ")
	}

	protoTour := mapper.MapToProtoTour(modelTour)
	response := &tour.TourResponse{
		Tour: protoTour,
	}
	return response, nil
}

func (handler *TourHandler) PostTour(ctx context.Context, request *tour.CreateTourRequest) (*tour.TourResponse, error) {
	var tourDto dto.TourDto

	newTour, createErr := handler.TourService.Save(&tourDto)
	if createErr != nil {
		println("Database exception: ")
	}

	protoTour := mapper.MapToProtoTour(newTour)
	response := &tour.TourResponse{
		Tour: protoTour,
	}
	return response, nil

}

func (handler *TourHandler) UpdateTour(ctx context.Context, request *tour.UpdateTourRequest) (*tour.TourResponse, error) {
	var tourDto dto.TourDto

	tourDto.Id = request.Id
	tourDto.BicycleTime = float64(request.Tour.BicycleTime)

	// Convert difficulty from string to int
	/*difficulty, err := tourDto.ParseDifficultyToInt()
	if err != nil {
		log.Println("Error while parsing difficulty:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update the difficulty field in tourDto
	tourDto.Difficulty = strconv.Itoa(difficulty)*/

	_, err1 := handler.TourService.Update(&tourDto)
	if err1 != nil {
		println("Database exception: ")
	}

	response := &tour.TourResponse{
		Tour: nil,
	}
	return response, nil
}

func (handler *TourHandler) DeleteTour(ctx context.Context, request *tour.GetRequest) (*tour.TourResponse, error) {
	id := request.Id
	log.Printf("Deleting tour with id: %s", id)

	handler.TourService.Delete(id)
	response := &tour.TourResponse{
		Tour: nil,
	}
	return response, nil
}

func (handler *TourHandler) GetToursByAuthor(ctx context.Context, request *tour.GetRequest) (*tour.GetToursResponse, error) {
	// Uzimanje ID autora iz URL parametra
	authorID := request.Id

	// Pretvaranje authorID u uint32
	id, err := strconv.ParseUint(authorID, 10, 32)
	if err != nil {
		log.Println("Error while parsing authorID:", err)

		return nil, err
	}
	convertedAuthorID := uint32(id)

	// Pozivanje servisa za dohvatanje tura za određenog autora
	tours, err := handler.TourService.GetToursByAuthor(convertedAuthorID)
	if err != nil {
		log.Println("Error while retrieving tours by author:", err)
		return nil, err
	}
	protoTours := mapper.MapSliceToProtoTours(tours)

	response := &tour.GetToursResponse{
		Tours: protoTours,
	}
	return response, nil
}

func (handler *TourHandler) GetCheckpoint(ctx context.Context, request *tour.GetRequest) (*tour.CheckpointResponse, error) {
	id := request.Id
	log.Printf("Checkpoint sa id-em %s", id)
	checkpoint, err := handler.CheckpointService.Get(id)
	if err != nil {
		println("Database exception: ")
	}

	protoCheckpoint := mapper.MapToProtoCheckpoint(checkpoint)
	response := &tour.CheckpointResponse{
		Checkpoint: protoCheckpoint,
	}
	return response, nil
}

func (handler *TourHandler) PostCheckpoint(ctx context.Context, request *tour.CreateCheckpointRequest) (*tour.CheckpointResponse, error) {

	checkpoint := mapper.MapToCheckpoint(request.Checkpoint)
	newCheckPoint, createErr := handler.CheckpointService.Save(checkpoint)

	if createErr != nil {
		println("Database exception: ")
	}

	protoCheckpoint := mapper.MapToProtoCheckpoint(newCheckPoint)
	response := &tour.CheckpointResponse{
		Checkpoint: protoCheckpoint,
	}
	return response, nil

}

func (handler *TourHandler) UpdateCheckpoint(ctx context.Context, request *tour.UpdateCheckpointRequest) (*tour.CheckpointResponse, error) {
	checkpoint := mapper.MapToCheckpoint(request.Checkpoint)

	handler.CheckpointService.Update(checkpoint)

	protoCheckpoint := mapper.MapToProtoCheckpoint(checkpoint)
	response := &tour.CheckpointResponse{
		Checkpoint: protoCheckpoint,
	}
	return response, nil

}

func (handler *TourHandler) DeleteCheckpoint(ctx context.Context, request *tour.GetRequest) (*tour.CheckpointResponse, error) {
	id := request.Id
	log.Printf("Deleting checkpoint with id: %s", id)

	handler.CheckpointService.Delete(id)
	response := &tour.CheckpointResponse{
		Checkpoint: nil,
	}
	return response, nil

}

/*func (handler *CheckpointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	checkpoints, err := handler.CheckpointService.GetAll()
	if err != nil {
		log.Println("Error while retrieving checkpoints:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(checkpoints)
}*/

func (handler *TourHandler) GetCheckpointsByTourId(ctx context.Context, request *tour.GetRequest) (*tour.GetCheckpointsResponse, error) {
	idStr := request.Id

	log.Printf("Get Checkpoints by tour id: %s", idStr)
	checkpoints, err := handler.CheckpointService.GetCheckpointsByTourID(idStr)
	if err != nil {
		log.Println("Error while retrieving tours by author:", err)
		return nil, err
	}
	protoCheckpoints := mapper.MapSliceToProtoCheckpoints(checkpoints)

	response := &tour.GetCheckpointsResponse{
		Checkpoints: protoCheckpoints,
	}
	return response, nil
}

func (handler *TourHandler) GetEquipment(ctx context.Context, request *tour.GetRequest) (*tour.EquipmentResponse, error) {
	id := request.Id
	log.Printf("Equipment sa id-em %s", id)
	equipment, err := handler.EquipmentService.Get(id)

	if err != nil {
		println("Database exception: ")
	}

	protoEquipment := mapper.MapToProtoEquipment(equipment)
	response := &tour.EquipmentResponse{
		Equipment: protoEquipment,
	}
	return response, nil

}

func (handler *TourHandler) PostEquipment(ctx context.Context, request *tour.CreateEquipmentRequest) (*tour.EquipmentResponse, error) {
	equipment := mapper.MapToEquipment(request.Equipment)

	err := handler.EquipmentService.Save(equipment)
	if err != nil {
		println("Database exception: ")
	}

	protoEquipment := mapper.MapToProtoEquipment(equipment)
	response := &tour.EquipmentResponse{
		Equipment: protoEquipment,
	}
	return response, nil
}

func (handler *TourHandler) UpdateEquipment(ctx context.Context, request *tour.UpdateEquipmentRequest) (*tour.EquipmentResponse, error) {
	equipment := mapper.MapToEquipment(request.Equipment)
	handler.EquipmentService.Update(equipment)

	protoEquipment := mapper.MapToProtoEquipment(equipment)
	response := &tour.EquipmentResponse{
		Equipment: protoEquipment,
	}
	return response, nil
}

func (handler *TourHandler) DeleteEquipment(ctx context.Context, request *tour.GetRequest) (*tour.EquipmentResponse, error) {
	id := request.Id
	log.Printf("Deleting equipment with id: %s", id)

	handler.EquipmentService.Delete(id)
	response := &tour.EquipmentResponse{
		Equipment: nil,
	}
	return response, nil

}

/*func (handler *EquipmentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	equipments, err := handler.EquipmentService.GetAll()
	if err != nil {
		log.Println("Error while retrieving equipments:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipments)
}*/

func (handler *TourHandler) GetTourProblem(ctx context.Context, request *tour.GetRequest) (*tour.TourProblemResponse, error) {
	id := request.Id
	log.Printf("TourProblem sa id-em %s", id)
	tourProblem, err := handler.TourProblemService.Get(id)
	if err != nil {
		println("Database exception: ")
	}

	protoTourProblem := mapper.MapToProtoTourProblem(tourProblem)
	response := &tour.TourProblemResponse{
		TourProblem: protoTourProblem,
	}
	return response, nil
}

func (handler *TourHandler) PostTourProblem(ctx context.Context, request *tour.CreateTourProblemRequest) (*tour.TourProblemResponse, error) {
	tourProblem := mapper.MapToTourProblem(request.TourProblem)

	err := handler.TourProblemService.Save(tourProblem)
	if err != nil {
		println("Database exception: ")
	}

	protoTourProblem := mapper.MapToProtoTourProblem(tourProblem)
	response := &tour.TourProblemResponse{
		TourProblem: protoTourProblem,
	}
	return response, nil
}

func (handler *TourHandler) UpdateTourProblem(ctx context.Context, request *tour.UpdateTourProblemRequest) (*tour.TourProblemResponse, error) {
	tourProblem := mapper.MapToTourProblem(request.TourProblem)

	handler.TourProblemService.Update(tourProblem)
	protoTourProblem := mapper.MapToProtoTourProblem(tourProblem)
	response := &tour.TourProblemResponse{
		TourProblem: protoTourProblem,
	}
	return response, nil
}

func (handler *TourHandler) DeleteTourProblem(ctx context.Context, request *tour.GetRequest) (*tour.TourProblemResponse, error) {
	id := request.Id
	log.Printf("Deleting TourProblem with id: %s", id)

	handler.TourProblemService.Delete(id)
	response := &tour.TourProblemResponse{
		TourProblem: nil,
	}
	return response, nil
}

/*func (handler *TourProblemHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tourProblems, err := handler.TourProblemService.GetAll()
	if err != nil {
		log.Println("Error while retrieving TourProblems:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourProblems)
}*/

func (handler *TourHandler) GetTourProblemsForTourist(ctx context.Context, request *tour.GetRequest) (*tour.GetTourProblemsResponse, error) {
	idStr := request.Id
	id, err := strconv.Atoi(idStr)
	log.Printf("Get TourProblem by tour id: %d", id)
	log.Printf("parsing ID to integer: %v", err)
	tourProblems, err := handler.TourProblemService.GetTourProblemsForTourist(uint32(id))
	if err != nil {
		log.Println("Error while retrieving tours by author:", err)
		return nil, err
	}
	protoTourProblems := mapper.MapSliceToProtoTourProblems(tourProblems)

	response := &tour.GetTourProblemsResponse{
		TourProblems: protoTourProblems,
	}
	return response, nil
}

/*func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("TourReview sa id-em %s", id)
	tourReview, err := handler.TourReviewService.Get(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourReview)
}*/

func (handler *TourHandler) PostTourReview(ctx context.Context, request *tour.CreateTourReviewRequest) (*tour.TourReviewResponse, error) {
	tourReview := mapper.MapToTourReview(request.TourReview)

	handler.TourReviewService.Save(tourReview)

	protoTourReview := mapper.MapToProtoTourReview(tourReview)
	response := &tour.TourReviewResponse{
		TourReview: protoTourReview,
	}
	return response, nil
}

func (handler *TourHandler) UpdateTourReview(ctx context.Context, request *tour.UpdateTourReviewRequest) (*tour.TourReviewResponse, error) {
	tourReview := mapper.MapToTourReview(request.TourReview)

	handler.TourReviewService.Update(tourReview)

	protoTourReview := mapper.MapToProtoTourReview(tourReview)
	response := &tour.TourReviewResponse{
		TourReview: protoTourReview,
	}
	return response, nil
}

func (handler *TourHandler) DeleteTourReview(ctx context.Context, request *tour.GetRequest) (*tour.TourReviewResponse, error) {
	id := request.Id
	log.Printf("Deleting TourReview with id: %s", id)

	handler.TourReviewService.Delete(id)
	response := &tour.TourReviewResponse{
		TourReview: nil,
	}
	return response, nil
}

func (handler *TourHandler) GetAllTourReviews(ctx context.Context, request *tour.GetAllRequest) (*tour.GetTourReviewsResponse, error) {
	tourReviews, err := handler.TourReviewService.GetAll()
	if err != nil {
		println("Database exception: ")
	}

	if tourReviews == nil {
		return nil, fmt.Errorf("failed to get blogs")
	}
	protoTourReviews := mapper.MapSliceToProtoTourReviews(tourReviews)
	response := &tour.GetTourReviewsResponse{
		TourReviews: protoTourReviews,
	}
	return response, nil
}

func (handler *TourHandler) GetTourReviewsByTourID(ctx context.Context, request *tour.GetRequest) (*tour.GetTourReviewsResponse, error) {
	idStr := request.Id

	log.Printf("Get tourReviews by tour id: %s", idStr)
	tourReviews, err := handler.TourReviewService.GetTourReviewsByTourID(idStr)
	if err != nil {
		log.Println("Error while retrieving tours by author:", err)
		return nil, err
	}
	protoTourReviews := mapper.MapSliceToProtoTourReviews(tourReviews)

	response := &tour.GetTourReviewsResponse{
		TourReviews: protoTourReviews,
	}
	return response, nil
}

func (handler *TourHandler) GetAverageGradeForTour(ctx context.Context, request *tour.GetRequest) (*tour.GetAverageGradeForTourRequest, error) {
	idStr := request.Id

	log.Printf("Get tourReviews by tour id: %s", idStr)
	averageGrade, err := handler.TourReviewService.GetAverageGradeForTour(idStr)
	if err != nil {
		log.Println("Error while retrieving tours by author:", err)
		return nil, err
	}

	//protoAverageGrades := mapper.MapToProtoAverageGrade(averageGrade)

	response := &tour.GetAverageGradeForTourRequest{
		AverageGrade: float32(averageGrade),
	}
	return response, nil

}
