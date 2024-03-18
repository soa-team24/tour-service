package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tour-service/model"
	"tour-service/service"

	"github.com/gorilla/mux"
)

type CheckpointHandler struct {
	CheckpointService *service.CheckpointService
}

func (handler *CheckpointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Checkpoint sa id-em %s", id)
	checkpoint, err := handler.CheckpointService.Get(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(checkpoint)
}

func (handler *CheckpointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var checkpoint model.Checkpoint
	err := json.NewDecoder(req.Body).Decode(&checkpoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CheckpointService.Save(&checkpoint)
	if err != nil {
		println("Error while creating a new checkpoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *CheckpointHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var checkpoint model.Checkpoint
	err := json.NewDecoder(req.Body).Decode(&checkpoint)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.CheckpointService.Update(&checkpoint)
	if err != nil {
		log.Println("Error while updating the checkpoint:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *CheckpointHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting checkpoint with id: %s", id)

	err := handler.CheckpointService.Delete(id)
	if err != nil {
		log.Println("Error while deleting the checkpoint:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *CheckpointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	checkpoints, err := handler.CheckpointService.GetAll()
	if err != nil {
		log.Println("Error while retrieving checkpoints:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(checkpoints)
}
