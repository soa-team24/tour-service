package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tour-service/model"
	"tour-service/service"

	"github.com/gorilla/mux"
)

type BundleHandler struct {
	BundleService *service.BundleService
}

func (handler *BundleHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Bundle sa id-em %s", id)
	bundle, err := handler.BundleService.Get(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(bundle)
}

func (handler *BundleHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var bundle model.Bundle
	err := json.NewDecoder(req.Body).Decode(&bundle)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BundleService.Save(&bundle)
	if err != nil {
		println("Error while creating a new bundle")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BundleHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var bundle model.Bundle
	err := json.NewDecoder(req.Body).Decode(&bundle)
	if err != nil {
		log.Println("Error while parsing JSON:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.BundleService.Update(&bundle)
	if err != nil {
		log.Println("Error while updating the bundle:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BundleHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Deleting bundle with id: %s", id)

	err := handler.BundleService.Delete(id)
	if err != nil {
		log.Println("Error while deleting the bundle:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *BundleHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	bundles, err := handler.BundleService.GetAll()
	if err != nil {
		log.Println("Error while retrieving bundles:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(bundles)
}
