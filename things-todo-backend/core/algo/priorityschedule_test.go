package algo_test

import (
	"testing"
	"time"

	"github.com/KKogaa/things-todo-backend/core/algo"
	"github.com/KKogaa/things-todo-backend/core/entity"
)

func TestPrioritySchedule(t *testing.T) {
	tasks := []*entity.Task{
		{Priority: entity.Low,
			Duration:   time.Duration(2 * time.Hour),
			Difficulty: entity.Hard,
		},
		{Priority: entity.Urgent,
			Duration:   time.Duration(2 * time.Hour),
			Difficulty: entity.Hard,
		},
		{Priority: entity.Urgent,
			Duration:   time.Duration(2 * time.Hour),
			Difficulty: entity.Hard,
		},
		{Priority: entity.Moderate,
			Duration:   time.Duration(2 * time.Hour),
			Difficulty: entity.Hard,
		},
	}

	taskManager := algo.NewTaskManager(tasks)

	prioritySchedule := algo.NewPrioritySchedule()
	prioritySchedule.Schedule(taskManager)

	if len(tasks) != len(taskManager.Tasks) {
		t.Errorf("Expected len of %d got %d", len(tasks),
			len(taskManager.Tasks))
	}

	for i := 0; i < len(taskManager.Tasks)-1; i++ {
		if taskManager.Tasks[i].Priority < taskManager.Tasks[i+1].Priority {
			t.Errorf("Expected %d index to be more than %d index", i, i+1)
		}
	}

}
