package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"soa/grpc/proto/tour"
	"tour-service/mapper"
	"tour-service/service"

	"github.com/gorilla/mux"
)

type TourReviewHandler struct {
	TourReviewService *service.TourReviewService
}

func (handler *TourReviewHandler) Get(writer http.ResponseWriter, req *http.Request) {
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
}

func (handler *TourReviewHandler) PostTourReview(ctx context.Context, request *tour.CreateTourReviewRequest) (*tour.TourReviewResponse, error) {
	tourReview := mapper.MapToTourReview(request.TourReview)

	handler.TourReviewService.Save(tourReview)

	protoTourReview := mapper.MapToProtoTourReview(tourReview)
	response := &tour.TourReviewResponse{
		TourReview: protoTourReview,
	}
	return response, nil
}

func (handler *TourReviewHandler) UpdateTourReview(ctx context.Context, request *tour.UpdateTourReviewRequest) (*tour.TourReviewResponse, error) {
	tourReview := mapper.MapToTourReview(request.TourReview)

	handler.TourReviewService.Update(tourReview)

	protoTourReview := mapper.MapToProtoTourReview(tourReview)
	response := &tour.TourReviewResponse{
		TourReview: protoTourReview,
	}
	return response, nil
}

func (handler *TourReviewHandler) DeleteTourReview(ctx context.Context, request *tour.GetRequest) (*tour.TourReviewResponse, error) {
	id := request.Id
	log.Printf("Deleting TourReview with id: %s", id)

	handler.TourReviewService.Delete(id)
	response := &tour.TourReviewResponse{
		TourReview: nil,
	}
	return response, nil
}

func (handler *TourReviewHandler) GetAllTourReviews(ctx context.Context, request *tour.GetAllRequest) (*tour.GetTourReviewsResponse, error) {
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

func (handler *TourReviewHandler) GetTourReviewsByTourID(ctx context.Context, request *tour.GetRequest) (*tour.GetTourReviewsResponse, error) {
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

func (handler *TourReviewHandler) GetAverageGradeForTour(ctx context.Context, request *tour.GetRequest) (*tour.GetAverageGradeForTourRequest, error) {
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
