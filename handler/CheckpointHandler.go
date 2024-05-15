package handler

import (
	"context"
	"log"
	"soa/grpc/proto/tour"
	"tour-service/mapper"
	"tour-service/service"
)

type CheckpointHandler struct {
	CheckpointService *service.CheckpointService
}

func (handler *CheckpointHandler) GetCheckpoint(ctx context.Context, request *tour.GetRequest) (*tour.CheckpointResponse, error) {
	id := request.Id
	log.Printf("Checkpoint sa id-em %s", id)
	checkpoint, err := handler.CheckpointService.Get(id)
	if err != nil {
		println("Database exception: ")
	}

	protoCheckpoint := mapper.MapToProtoCheckpoint(checkpoint)
	response := &tour.CheckpointResponse{
		Checkpoint: protoCheckpoint,
	}
	return response, nil
}

func (handler *CheckpointHandler) PostCheckpoint(ctx context.Context, request *tour.CreateCheckpointRequest) (*tour.CheckpointResponse, error) {

	checkpoint := mapper.MapToCheckpoint(request.Checkpoint)
	newCheckPoint, createErr := handler.CheckpointService.Save(checkpoint)

	if createErr != nil {
		println("Database exception: ")
	}

	protoCheckpoint := mapper.MapToProtoCheckpoint(newCheckPoint)
	response := &tour.CheckpointResponse{
		Checkpoint: protoCheckpoint,
	}
	return response, nil

}

func (handler *CheckpointHandler) UpdateCheckpoint(ctx context.Context, request *tour.UpdateCheckpointRequest) (*tour.CheckpointResponse, error) {
	checkpoint := mapper.MapToCheckpoint(request.Checkpoint)

	handler.CheckpointService.Update(checkpoint)

	protoCheckpoint := mapper.MapToProtoCheckpoint(checkpoint)
	response := &tour.CheckpointResponse{
		Checkpoint: protoCheckpoint,
	}
	return response, nil

}

func (handler *CheckpointHandler) DeleteCheckpoint(ctx context.Context, request *tour.GetRequest) (*tour.CheckpointResponse, error) {
	id := request.Id
	log.Printf("Deleting checkpoint with id: %s", id)

	handler.CheckpointService.Delete(id)
	response := &tour.CheckpointResponse{
		Checkpoint: nil,
	}
	return response, nil

}

/*func (handler *CheckpointHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	checkpoints, err := handler.CheckpointService.GetAll()
	if err != nil {
		log.Println("Error while retrieving checkpoints:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(checkpoints)
}*/

func (handler *CheckpointHandler) GetCheckpointsByTourId(ctx context.Context, request *tour.GetRequest) (*tour.GetCheckpointsResponse, error) {
	idStr := request.Id

	log.Printf("Get Checkpoints by tour id: %s", idStr)
	checkpoints, err := handler.CheckpointService.GetCheckpointsByTourID(idStr)
	if err != nil {
		log.Println("Error while retrieving tours by author:", err)
		return nil, err
	}
	protoCheckpoints := mapper.MapSliceToProtoCheckpoints(checkpoints)

	response := &tour.GetCheckpointsResponse{
		Checkpoints: protoCheckpoints,
	}
	return response, nil
}
