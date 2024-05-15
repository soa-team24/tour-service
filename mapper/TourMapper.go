package mapper

import (
	"tour-service/model"

	"soa/grpc/proto/tour"
)

func MapSliceToProtoToursPointer(modelTours []*model.Tour) []*tour.Tour {
	protoTours := make([]*tour.Tour, 0, len(modelTours))

	for _, modelTour := range modelTours {
		protoTour := MapToProtoTour(modelTour)
		protoTours = append(protoTours, protoTour)
	}

	return protoTours
}
func MapSliceToProtoTours(modelTours []model.Tour) []*tour.Tour {
	var protoTours []*tour.Tour

	for _, modelTour := range modelTours {
		protoTour := MapToProtoTour(&modelTour)
		protoTours = append(protoTours, protoTour)
	}

	return protoTours
}

func MapToProtoTour(modelTour *model.Tour) *tour.Tour {
	protoTour := &tour.Tour{
		Id:          modelTour.ID.String(),
		AuthorId:    modelTour.AuthorID,
		Name:        modelTour.Name,
		Description: modelTour.Description,
		PublishTime: modelTour.PublishTime,
		Status:      tour.Tour_Status(modelTour.Status),
		Image:       modelTour.Image,
		Difficulty:  uint32(modelTour.Difficulty),
		Price:       float32(modelTour.Price),
		FootTime:    float32(modelTour.FootTime),
		BicycleTime: float32(modelTour.BicycleTime),
		CarTime:     float32(modelTour.CarTime),
		TotalLength: float32(modelTour.TotalLength),
	}

	return protoTour
}
