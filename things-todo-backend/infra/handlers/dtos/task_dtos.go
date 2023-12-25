package dtos

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int8   `json:"duration"`
	Priority    int8   `json:"priority"`
	Difficulty  int8   `json:"difficulty"`
}
