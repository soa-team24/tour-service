package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tour-service/dto"
	"tour-service/model"
	"tour-service/service"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Tour sa id-em %s", id)
	tour, err := handler.TourService.Get(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	tourDto := &dto.TourDto{
		Name:        tour.Title,
		Description: tour.Description,
		PublishTime: tour.PublishTime,
		Status:      int(tour.Status),
		Image:       tour.Image,
		Difficulty:  strconv.Itoa(tour.Difficulty),
		Price:       tour.Price,
		FootTime:    tour.FootTime,
		BicycleTime: tour.BicycleTime,
		CarTime:     tour.CarTime,
		TotalLength: tour.TotalLength,
		AuthorID:    tour.AuthorID,
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourDto)
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourDto dto.TourDto
	err := json.NewDecoder(req.Body).Decode(&tourDto)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	newTour, createErr := handler.TourService.Save(&tourDto)
	if createErr != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(newTour)
}

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.Update(&tour)
	if err != nil {
		log.Println("Error while updating the tour:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting tour with id: %s", id)

	err := handler.TourService.Delete(id)
	if err != nil {
		log.Println("Error while deleting the tour:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *TourHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tours, err := handler.TourService.GetAll()
	if err != nil {
		log.Println("Error while retrieving tours:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tours)
}
