package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tour-service/model"
	"tour-service/service"

	"github.com/gorilla/mux"
)

type TourProblemHandler struct {
	TourProblemService *service.TourProblemService
}

func (handler *TourProblemHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("TourProblem sa id-em %s", id)
	tourProblem, err := handler.TourProblemService.Get(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourProblem)
}

func (handler *TourProblemHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourProblem model.TourProblem
	err := json.NewDecoder(req.Body).Decode(&tourProblem)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourProblemService.Save(&tourProblem)
	if err != nil {
		println("Error while creating a new TourProblem")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourProblemHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tourProblem model.TourProblem
	err := json.NewDecoder(req.Body).Decode(&tourProblem)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourProblemService.Update(&tourProblem)
	if err != nil {
		log.Println("Error while updating the TourProblem:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourProblemHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting TourProblem with id: %s", id)

	err := handler.TourProblemService.Delete(id)
	if err != nil {
		log.Println("Error while deleting the TourProblem:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *TourProblemHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	tourProblems, err := handler.TourProblemService.GetAll()
	if err != nil {
		log.Println("Error while retrieving TourProblems:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourProblems)
}
