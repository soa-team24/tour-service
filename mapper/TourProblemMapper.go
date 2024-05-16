package mapper

import (
	"time"
	"tour-service/model"

	"tour-service/proto/tour"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapSliceToProtoTourProblems(modelTourProblems []model.TourProblem) []*tour.TourProblem {
	var protoTourProblems []*tour.TourProblem

	for _, modelTourProblem := range modelTourProblems {
		protoTourProblem := MapToProtoTourProblem(&modelTourProblem)
		protoTourProblems = append(protoTourProblems, protoTourProblem)
	}

	return protoTourProblems
}

func MapToProtoTourProblem(modelTourProblem *model.TourProblem) *tour.TourProblem {
	protoTourProblem := &tour.TourProblem{
		Id:                modelTourProblem.ID.String(),
		ProblemCategory:   modelTourProblem.ProblemCategory,
		ProblemPriority:   modelTourProblem.ProblemPriority,
		Description:       modelTourProblem.Description,
		TimeStamp:         timestamppb.New(modelTourProblem.TimeStamp),
		TourId:            modelTourProblem.TourId,
		IsClosed:          modelTourProblem.IsClosed,
		IsResolved:        modelTourProblem.IsResolved,
		TouristId:         modelTourProblem.TouristId,
		DeadlineTimeStamp: timestamppb.New(*modelTourProblem.DeadlineTimeStamp),
	}

	return protoTourProblem
}

func MapToTourProblem(tourProblemP *tour.TourProblem) *model.TourProblem {
	var deadlineTime *time.Time
	if tourProblemP.DeadlineTimeStamp != nil {
		tempTime := tourProblemP.DeadlineTimeStamp.AsTime()
		deadlineTime = &tempTime
	}
	tourProblem := &model.TourProblem{

		ProblemCategory:   tourProblemP.ProblemCategory,
		ProblemPriority:   tourProblemP.ProblemPriority,
		Description:       tourProblemP.Description,
		TimeStamp:         tourProblemP.TimeStamp.AsTime(),
		TourId:            tourProblemP.TourId,
		IsClosed:          tourProblemP.IsClosed,
		IsResolved:        tourProblemP.IsResolved,
		TouristId:         tourProblemP.TouristId,
		DeadlineTimeStamp: deadlineTime,
	}

	return tourProblem

}
