package mapper

import (
	"tour-service/model"
	"tour-service/proto/tour"
)

func MapSliceToProtoCheckpoints(modelCheckpoints []model.Checkpoint) []*tour.Checkpoint {
	var protoCheckpoints []*tour.Checkpoint

	for _, modelCheckpoint := range modelCheckpoints {
		protoCheckpoint := MapToProtoCheckpoint(&modelCheckpoint)
		protoCheckpoints = append(protoCheckpoints, protoCheckpoint)
	}

	return protoCheckpoints
}

func MapToProtoCheckpoint(modelCheckpoint *model.Checkpoint) *tour.Checkpoint {
	protoCheckpoint := &tour.Checkpoint{
		Id:          modelCheckpoint.ID.String(),
		TourId:      modelCheckpoint.TourID,
		Latitude:    float32(modelCheckpoint.Latitude),
		Longitude:   float32(modelCheckpoint.Longitude),
		Name:        modelCheckpoint.Name,
		Description: modelCheckpoint.Description,
		Image:       modelCheckpoint.Image,
		IsPublic:    modelCheckpoint.IsPublic,
	}

	return protoCheckpoint
}

func MapToCheckpoint(checkpointP *tour.Checkpoint) *model.Checkpoint {
	checkpoint := &model.Checkpoint{

		TourID:      checkpointP.TourId,
		Latitude:    float64(checkpointP.Latitude),
		Longitude:   float64(checkpointP.Longitude),
		Name:        checkpointP.Name,
		Description: checkpointP.Description,
		Image:       checkpointP.Image,
		IsPublic:    checkpointP.IsPublic,
	}

	return checkpoint

}
