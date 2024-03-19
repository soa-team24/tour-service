package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tour-service/model"
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

func (handler *TourReviewHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourReview model.TourReview
	err := json.NewDecoder(req.Body).Decode(&tourReview)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourReviewService.Save(&tourReview)
	if err != nil {
		println("Error while creating a new TourReview")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourReviewHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tourReview model.TourReview
	err := json.NewDecoder(req.Body).Decode(&tourReview)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourReviewService.Update(&tourReview)
	if err != nil {
		log.Println("Error while updating the TourReview:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourReviewHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting TourReview with id: %s", id)

	err := handler.TourReviewService.Delete(id)
	if err != nil {
		log.Println("Error while deleting the TourReview:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *TourReviewHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tourReviews, err := handler.TourReviewService.GetAll()
	if err != nil {
		log.Println("Error while retrieving tourReviews:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourReviews)
}
