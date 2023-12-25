package entity

import "time"

// todo move Difficulty to enum
// todo move Duration to enum
const (
	Easy   Difficulty = 1
	Medium            = 2
	Hard              = 3
)

type Difficulty int8

const (
	Low      Priority = 1
	Moderate          = 2
	Urgent            = 3
)

type Priority int8

type Task struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Difficulty  Difficulty    `json:"difficulty"`
	Duration    time.Duration `json:"duration"`
	Priority    Priority      `json:"priority"`
}

func NewTask(duration time.Duration, difficulty Difficulty,
	priority Priority) Task {

	return Task{
		Duration:   duration,
		Difficulty: difficulty,
		Priority:   AssignPriority(priority),
	}

}

func AssignPriority(priority Priority) Priority {
	//manually assign priority or assign priority using algo basing on other tasks
	if priority != 0 {
		return priority
	}

	return priority

}
