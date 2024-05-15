package mapper

import (
	"tour-service/model"

	"soa/grpc/proto/tour"
)

func MapSliceToProtoEquipments(modelEquipments []model.Equipment) []*tour.Equipment {
	var protoEquipments []*tour.Equipment

	for _, modelEquipment := range modelEquipments {
		protoEquipment := MapToProtoEquipment(&modelEquipment)
		protoEquipments = append(protoEquipments, protoEquipment)
	}

	return protoEquipments
}

func MapToProtoEquipment(modelEquipment *model.Equipment) *tour.Equipment {
	protoEquipment := &tour.Equipment{
		Id:          modelEquipment.ID.String(),
		Name:        modelEquipment.Name,
		Description: modelEquipment.Description,
	}

	return protoEquipment
}

func MapToEquipment(equipmentP *tour.Equipment) *model.Equipment {
	equipment := &model.Equipment{

		Name:        equipmentP.Name,
		Description: equipmentP.Description,
	}

	return equipment

}
