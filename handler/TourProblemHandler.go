package handler

import (
	"context"
	"log"
	"strconv"
	"tour-service/mapper"
	"tour-service/proto/tour"
	"tour-service/service"
)

type TourProblemHandler struct {
	TourProblemService *service.TourProblemService
}

func (handler *TourProblemHandler) GetTourProblem(ctx context.Context, request *tour.GetRequest) (*tour.TourProblemResponse, error) {
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

func (handler *TourProblemHandler) PostTourProblem(ctx context.Context, request *tour.CreateTourProblemRequest) (*tour.TourProblemResponse, error) {
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

func (handler *TourProblemHandler) UpdateTourProblem(ctx context.Context, request *tour.UpdateTourProblemRequest) (*tour.TourProblemResponse, error) {
	tourProblem := mapper.MapToTourProblem(request.TourProblem)

	handler.TourProblemService.Update(tourProblem)
	protoTourProblem := mapper.MapToProtoTourProblem(tourProblem)
	response := &tour.TourProblemResponse{
		TourProblem: protoTourProblem,
	}
	return response, nil
}

func (handler *TourProblemHandler) DeleteTourProblem(ctx context.Context, request *tour.GetRequest) (*tour.TourProblemResponse, error) {
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

func (handler *TourProblemHandler) GetTourProblemsForTourist(ctx context.Context, request *tour.GetRequest) (*tour.GetTourProblemsResponse, error) {
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
