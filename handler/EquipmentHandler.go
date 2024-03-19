package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tour-service/model"
	"tour-service/service"

	"github.com/gorilla/mux"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (handler *EquipmentHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Equipment sa id-em %s", id)
	equipment, err := handler.EquipmentService.Get(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EquipmentService.Save(&equipment)
	if err != nil {
		println("Error while creating a new equipment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EquipmentHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EquipmentService.Update(&equipment)
	if err != nil {
		log.Println("Error while updating the equipment:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EquipmentHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting equipment with id: %s", id)

	err := handler.EquipmentService.Delete(id)
	if err != nil {
		log.Println("Error while deleting the equipment:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *EquipmentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	equipments, err := handler.EquipmentService.GetAll()
	if err != nil {
		log.Println("Error while retrieving equipments:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipments)
}
