package handler

import (
	"context"
	"log"
	"tour-service/mapper"
	"tour-service/proto/tour"
	"tour-service/service"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (handler *EquipmentHandler) GetEquipment(ctx context.Context, request *tour.GetRequest) (*tour.EquipmentResponse, error) {
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

func (handler *EquipmentHandler) PostEquipment(ctx context.Context, request *tour.CreateEquipmentRequest) (*tour.EquipmentResponse, error) {
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

func (handler *EquipmentHandler) UpdateEquipment(ctx context.Context, request *tour.UpdateEquipmentRequest) (*tour.EquipmentResponse, error) {
	equipment := mapper.MapToEquipment(request.Equipment)
	handler.EquipmentService.Update(equipment)

	protoEquipment := mapper.MapToProtoEquipment(equipment)
	response := &tour.EquipmentResponse{
		Equipment: protoEquipment,
	}
	return response, nil
}

func (handler *EquipmentHandler) DeleteEquipment(ctx context.Context, request *tour.GetRequest) (*tour.EquipmentResponse, error) {
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
