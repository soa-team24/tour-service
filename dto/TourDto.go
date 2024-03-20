package dto

type TourDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
	Difficulty  string  `json:"difficulty"`
	TotalLength float64 `json:"totalLength"`
	FootTime    float64 `json:"footTime"`
	BicycleTime float64 `json:"bicycleTime"`
	CarTime     float64 `json:"carTime"`
	AuthorID    uint32  `json:"authorId"`
	PublishTime string  `json:"publishTime"`
	Price       float64 `json:"price"`
	Points      int     `json:"points"`
	Image       string  `json:"image"`
}
