package mapper

import (
	"tour-service/model"

	"tour-service/proto/tour"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapSliceToProtoTourReviews(modelTourReviews []model.TourReview) []*tour.TourReview {
	var protoTourReviews []*tour.TourReview

	for _, modelTourReview := range modelTourReviews {
		protoTourReview := MapToProtoTourReview(&modelTourReview)
		protoTourReviews = append(protoTourReviews, protoTourReview)
	}

	return protoTourReviews
}

func MapToProtoTourReview(modelTourReview *model.TourReview) *tour.TourReview {
	protoTourReview := &tour.TourReview{
		Id:         modelTourReview.ID.String(),
		Grade:      modelTourReview.Grade,
		Comment:    modelTourReview.Comment,
		UserId:     modelTourReview.UserID,
		VisitDate:  timestamppb.New(modelTourReview.VisitDate),
		ReviewDate: timestamppb.New(modelTourReview.ReviewDate),
		Images:     modelTourReview.Images,
		TourId:     modelTourReview.TourID,
	}

	return protoTourReview
}

func MapToTourReview(tourReviewP *tour.TourReview) *model.TourReview {
	tourReview := &model.TourReview{

		Grade:      tourReviewP.Grade,
		Comment:    tourReviewP.Comment,
		UserID:     tourReviewP.UserId,
		VisitDate:  tourReviewP.VisitDate.AsTime(),
		ReviewDate: tourReviewP.ReviewDate.AsTime(),
		Images:     tourReviewP.Images,
		TourID:     tourReviewP.TourId,
	}

	return tourReview

}
