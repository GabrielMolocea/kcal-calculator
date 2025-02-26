package types

type Food struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Calories int64  `json:"calories"`
}

type TotalCalories struct {
	Total int64 `json:"total"`
}

type TargetKcal struct {
	ID         int64 `json:"id"`
	TargetKcal int64 `json:"targetKcal"`
}